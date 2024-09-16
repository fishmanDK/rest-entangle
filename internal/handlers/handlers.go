package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fishmanDK/rest-entangle/internal/service"
)

type Handlers struct {
	service *service.Service
}


func New(_service *service.Service) *Handlers{
	return &Handlers{
		service: _service,
	}
}

func (h *Handlers) Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	amount, err := h.service.GetAmount()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"amount": amount})
}

