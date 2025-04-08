package main

import (
	"log"
	"net/http"

	"github.com/dhruvsaxena1998/aio/cmd/internal/routes"
)

func main() {
	r := routes.New()
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
