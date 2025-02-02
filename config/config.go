package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server Server
	PG     PG
}

type Server struct {
	Port string
}

type PG struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	SSLMODE  string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (a *AppConfig) Init() error {
	if err := godotenv.Load(); err != nil {
		return errors.New("Coudn't load env file: " + err.Error())
	}

	a.Server.Port = os.Getenv("PORT")

	a.PG.Host = os.Getenv("POSTGRES_HOST")
	a.PG.Port = os.Getenv("POSTGRES_PORT")
	a.PG.User = os.Getenv("POSTGRES_USER")
	a.PG.Password = os.Getenv("POSTGRES_PASSWORD")
	a.PG.DB = os.Getenv("POSTGRES_DB")
	a.PG.SSLMODE = os.Getenv("POSTGRES_SSLMODE")

	return nil
}
