package main

import (
	"log"
	"os"

	"github.com/ramisoul84/reset-server/config"
	"github.com/ramisoul84/reset-server/internal/server/http"
	"github.com/ramisoul84/reset-server/internal/service/domain"
	"github.com/ramisoul84/reset-server/internal/storage/postgres"
)

func main() {
	// App Configuration
	cfg := config.NewAppConfig()
	err := cfg.Init()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Println("App Configuration Initialized Successfully")

	// Connect to Postgres
	db, err := postgres.New(cfg.PG)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Println("Successfully connected to the PostgreSQL database!", db)

	// repositories
	repositories := postgres.NewRepository(db)
	// services
	service := domain.NewService(repositories)

	// Run a server
	srv := http.NewServer(cfg, service)
	err = srv.Serve()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

}
