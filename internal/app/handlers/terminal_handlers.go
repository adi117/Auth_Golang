package handlers

import (
	"auth/internal/app/models"
	"auth/internal/app/services"
	"encoding/json"
	"net/http"
)

type TerminalHandlers struct {
	TerminalService services.TerminalServiceInterface
}

func NewTerminalHandlers(service services.TerminalServiceInterface) *TerminalHandlers {
	return &TerminalHandlers{TerminalService: service}
}

func (h *TerminalHandlers) CreateTerminal(w http.ResponseWriter, r *http.Request) {
	var req models.TerminalRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	terminal, err := h.TerminalService.CreateTerminal(req.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(terminal)
}
