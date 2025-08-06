package routes

import (
	"auth/internal/app/handlers"
	"auth/internal/middleware"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux, handler *handlers.AuthHandlers) {
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/protected", middleware.JWTMiddleware(middleware.ProtectedEndpoint))
}
