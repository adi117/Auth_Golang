package handlers

import (
	"auth/internal/app/models"
	"auth/internal/app/services"
	"auth/internal/utils"
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

	user.Password = "" // set blank password in response

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

	token, err := utils.GenerateJWT(user.Id, user.Username)

	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	resp := struct {
		User  models.UserReponse   `json:"user"`
		Token *utils.TokenResponse `json:"token"`
	}{
		User: models.UserReponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		},
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
