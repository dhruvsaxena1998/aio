package middlewares

import (
	"errors"
	"net/http"

	"context"

	"github.com/dhruvsaxena1998/aio/cmd/internal/database"
	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
	"gorm.io/gorm"
)

type ContextKey string

const (
	UserContextKey ContextKey = "user"
)

func RequireAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "API Key is required", http.StatusUnauthorized)
			return
		}

		var user models.User
		err := database.DB.
			Where(&models.User{APIKey: apiKey}).
			First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireSuperAdministrator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := GetUserFromContext(r.Context())
		if !ok {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}

		hasPermission, err := user.HasPermissions(database.DB, "allow:all")
		if err != nil || !hasPermission {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetUserFromContext(ctx context.Context) (*models.User, bool) {
	if user, ok := ctx.Value(UserContextKey).(*models.User); ok {
		return user, true
	}
	return nil, false
}
