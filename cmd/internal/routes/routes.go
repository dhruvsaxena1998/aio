package routes

import (
	"net/http"

	"github.com/dhruvsaxena1998/aio/cmd/internal/handlers"
	"github.com/dhruvsaxena1998/aio/cmd/internal/middlewares"
	"github.com/go-chi/chi/v5"
	m "github.com/go-chi/chi/v5/middleware"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Use(m.RequestID)   // Adds X-Request-ID header
	r.Use(m.RealIP)      // Gets client IP from headers
	r.Use(m.Logger)      // Logs the start and end of each request
	r.Use(m.Recoverer)   // Recovers from panics and logs
	r.Use(m.Timeout(60)) // Times out requests after 60s

	r.Get("/ping", handlers.PingHandler)

	// Admin only routes
	r.Route("/api/v1", func(api chi.Router) {
		api.Use(middlewares.OnlyAdmin)

		api.Get("/secret", handlers.SecretHandler)
	})

	return r
}
