package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ramisoul84/reset-server/config"
)

func New(cfg config.PG) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", cfg.User, cfg.Password, cfg.DB, cfg.Host, cfg.Port, cfg.SSLMODE)
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err = CreateTables(db); err != nil {
		return nil, err
	}
	return db, nil

}
