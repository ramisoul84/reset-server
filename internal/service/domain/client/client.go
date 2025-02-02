package client

import (
	"errors"
	"fmt"

	"github.com/ramisoul84/reset-server/internal/entity/client"
	"github.com/ramisoul84/reset-server/internal/storage"
)

type ClientService struct {
	clientRepository storage.ClientRepository
}

func NewClientService(clientRepository storage.ClientRepository) *ClientService {
	return &ClientService{
		clientRepository,
	}
}

func (s *ClientService) GetClientByEmail(clientEmail string) (*client.Client, error) {
	return s.clientRepository.GetByEmail(clientEmail)
}

func (s *ClientService) CreateClient(client client.Client) (*client.Client, error) {
	foundClient, _ := s.clientRepository.GetByEmail(client.Email)
	fmt.Println("service", foundClient)
	if foundClient != nil {
		return nil, errors.New("email already registered")
	}
	return s.clientRepository.Add(client)
}

func (s *ClientService) ListClients() ([]*client.Client, error) {
	return s.clientRepository.List()
}
