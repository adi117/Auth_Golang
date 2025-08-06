package main

import (
	"auth/internal/app/handlers"
	"auth/internal/app/repositories"
	"auth/internal/app/services"
	"auth/internal/config"
	"auth/internal/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	mux := http.NewServeMux()

	if err != nil {
		log.Fatal("Error to load .env file!")
	}
	setting := config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db := config.InitDB(setting)

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandlers := handlers.NewAuthHandlers(authService)

	terminalRepo := repositories.NewTerminalRepository(db)
	terminalService := services.NewTerminalService(terminalRepo)
	terminalHandlers := handlers.NewTerminalHandlers(terminalService)

	routes.RegisterAuthRoutes(mux, authHandlers)
	routes.RegisterTerminalRoutes(mux, terminalHandlers)

	http.ListenAndServe(":8080", mux)
}
