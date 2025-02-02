package storage

import (
	"github.com/ramisoul84/reset-server/internal/entity/client"
)

type ClientRepository interface {
	GetByEmail(clientEmail string) (*client.Client, error)
	Add(client client.Client) error
	List() ([]*client.Client, error)
}

type Repository struct {
	Client ClientRepository
}
