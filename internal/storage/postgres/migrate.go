package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateTables(db *sqlx.DB) error {
	createClientsQuery := `
	CREATE TABLE IF NOT EXISTS clients (
    	id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		email VARCHAR(255) NOT NULL UNIQUE
		);
	`
	_, err := db.Exec(createClientsQuery)
	if err != nil {
		log.Println("error creating clients table")
		return err
	}
	log.Println("Database migrated successfully")
	return nil
}
