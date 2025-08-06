package routes

import (
	"auth/internal/app/handlers"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux, handler *handlers.AuthHandlers) {
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)
}
