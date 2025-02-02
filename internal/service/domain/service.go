package domain

import (
	"github.com/ramisoul84/reset-server/internal/service"
	"github.com/ramisoul84/reset-server/internal/service/domain/client"
	"github.com/ramisoul84/reset-server/internal/storage"
)

func NewService(repositories *storage.Repository) *service.Service {
	return &service.Service{
		ClientService: client.NewClientService(repositories.Client),
	}
}
