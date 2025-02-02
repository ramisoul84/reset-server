package http

import (
	"net/http"

	"github.com/ramisoul84/reset-server/config"
	"github.com/ramisoul84/reset-server/internal/service"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.AppConfig, service *service.Service) *Server {
	handler := NewRouter(service)
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Server.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Serve() error {
	return s.httpServer.ListenAndServe()
}
