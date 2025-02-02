package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ramisoul84/reset-server/internal/server/http/handler/client"
	"github.com/ramisoul84/reset-server/internal/service"
)

func NewRouter(service *service.Service) http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Mount("/client", client.RegisterRoutes(service))

	})
	return r
}
