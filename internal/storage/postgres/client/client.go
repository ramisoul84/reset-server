package client

import (
	"github.com/jmoiron/sqlx"
	"github.com/ramisoul84/reset-server/internal/entity/client"
)

type ClientRepository struct {
	db *sqlx.DB
}

func NewClientRepository(db *sqlx.DB) *ClientRepository {
	return &ClientRepository{
		db,
	}
}

func (r *ClientRepository) GetByEmail(clientEmail string) (*client.Client, error) {
	var client client.Client
	query := `SELECT id, name, email, created_at FROM clients WHERE email = $1`
	err := r.db.Get(&client, query, clientEmail)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) Add(client client.Client) error {
	query := `
        INSERT INTO clients (name, email,created_at)
        VALUES (:name, :email, :created_at)`
	_, err := r.db.NamedExec(query, client)
	if err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) List() ([]*client.Client, error) {
	var clients []*client.Client
	query := `SELECT * FROM clients`
	err := r.db.Select(&clients, query)
	if err != nil {
		return nil, err
	}
	return clients, nil
}
