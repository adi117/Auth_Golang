package handlers

import (
	"auth/internal/app/models"
	"auth/internal/app/services"
	"encoding/json"
	"net/http"
)

type AuthHandlers struct {
	AuthService services.AuthServiceInterface
}

func NewAuthHandlers(service services.AuthServiceInterface) *AuthHandlers {
	return &AuthHandlers{AuthService: service}
}

func (h *AuthHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var req models.UserRegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.AuthService.Register(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var req models.UserLoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.AuthService.Login(req)

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user)
}
