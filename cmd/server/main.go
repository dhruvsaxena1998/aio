package main

import (
	"log"
	"net/http"

	"github.com/dhruvsaxena1998/aio/cmd/internal/database"
	"github.com/dhruvsaxena1998/aio/cmd/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Proceeding with defaults...")
	}

	database.Init()

	r := routes.New()
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
