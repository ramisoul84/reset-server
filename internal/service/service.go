package service

import (
	"github.com/ramisoul84/reset-server/internal/entity/client"
)

type ClientService interface {
	GetClientByEmail(clientEmail string) (*client.Client, error)
	CreateClient(client client.Client) (*client.Client, error)
	ListClients() ([]*client.Client, error)
}

type Service struct {
	ClientService ClientService
}
