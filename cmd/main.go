package main

import (
	"auth/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file!")
	}
	config.InitDB()
}
