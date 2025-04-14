package routes

import (
	"net/http"

	"github.com/dhruvsaxena1998/aio/cmd/internal/handlers"
	"github.com/dhruvsaxena1998/aio/cmd/internal/handlers/habits"
	"github.com/dhruvsaxena1998/aio/cmd/internal/middlewares"
	"github.com/go-chi/chi/v5"
	m "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(m.RequestID)   // Adds X-Request-ID header
	r.Use(m.RealIP)      // Gets client IP from headers
	r.Use(m.Logger)      // Logs the start and end of each request
	r.Use(m.Recoverer)   // Recovers from panics and logs
	r.Use(m.Timeout(60)) // Times out requests after 60s
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "HEAD"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-API-KEY"},
	}))

	r.Get("/ping", handlers.PingHandler)

	// Admin only routes
	r.Route("/api/v1", func(api chi.Router) {
		m := middlewares.NewMiddleware(db)

		api.Use(m.RequireAPIKey)

		api.With(m.RequireSuperAdministrator).
			Get("/secret", handlers.SecretHandler)

			// Habit routes
		api.Route("/habits", func(r chi.Router) {
			h := habits.NewHandler(db)

			r.Get("/", h.GetHabitsHandler)
			r.Post("/", h.CreateHabitHandler)

			r.Get("/{habitId}/completions", h.GetCompletionsHandler)
			r.Post("/{habitId}/completions", h.CreateCompletionHandler)
		})

	})

	return r
}
