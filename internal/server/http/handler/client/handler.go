package client

import (
	"encoding/json"
	"net/http"

	"github.com/ramisoul84/reset-server/internal/entity/client"
	"github.com/ramisoul84/reset-server/internal/service"
)

type httpHandler struct {
	service *service.Service
}

func newHTTPHandler(service *service.Service) *httpHandler {
	return &httpHandler{
		service,
	}
}

func (h *httpHandler) addClient(w http.ResponseWriter, r *http.Request) {
	var client client.Client
	json.NewDecoder(r.Body).Decode(&client)
	if len(client.Email) < 3 {
		http.Error(w, "wrong email", http.StatusBadRequest)
		return
	}
	addedClient, err := h.service.ClientService.CreateClient(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if err = json.NewEncoder(w).Encode(addedClient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *httpHandler) listClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.service.ClientService.ListClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if err = json.NewEncoder(w).Encode(clients); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
