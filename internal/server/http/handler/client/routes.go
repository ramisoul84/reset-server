package client

import (
	"github.com/go-chi/chi"
	"github.com/ramisoul84/reset-server/internal/service"
)

func RegisterRoutes(service *service.Service) *chi.Mux {
	r := chi.NewRouter()
	handler := newHTTPHandler(service)
	r.Get("/list", handler.listClients)
	r.Post("/register", handler.addClient)
	return r
}
