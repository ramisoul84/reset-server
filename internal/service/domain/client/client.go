package client

import (
	"errors"
	"time"

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

func (s *ClientService) CreateClient(client client.Client) error {
	foundClient, _ := s.clientRepository.GetByEmail(client.Email)
	if foundClient != nil {
		return errors.New("email already registered")
	}
	client.CreatedAt = time.Now()
	return s.clientRepository.Add(client)
}

func (s *ClientService) ListClients() ([]*client.Client, error) {
	return s.clientRepository.List()
}
