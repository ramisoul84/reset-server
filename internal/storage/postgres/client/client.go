package client

import (
	"errors"
	"fmt"

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
	client := &client.Client{}
	err := r.db.QueryRow("SELECT * FROM clients WHERE email = $1", clientEmail).Scan(&client.ID, &client.Name, &client.Email)
	fmt.Println("repo ", client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *ClientRepository) Add(client client.Client) (*client.Client, error) {
	query := "INSERT INTO clients (name,email) VALUES ($1,$2) RETURNING *"
	err := r.db.QueryRow(query, client.Name, client.Email).Scan(&client.ID, &client.Name, &client.Email)
	if err != nil {
		return nil, errors.New("Error creating user: " + err.Error())
	}
	return &client, nil
}

func (r *ClientRepository) List() ([]*client.Client, error) {
	clients := []*client.Client{}
	err := r.db.Select(&clients, "SELECT * FROM clients ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	return clients, nil
}
