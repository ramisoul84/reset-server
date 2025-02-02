package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/ramisoul84/reset-server/internal/storage"
	"github.com/ramisoul84/reset-server/internal/storage/postgres/client"
)

func NewRepository(db *sqlx.DB) *storage.Repository {
	return &storage.Repository{
		Client: client.NewClientRepository(db),
	}
}
