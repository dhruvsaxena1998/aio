package middlewares

import (
	"errors"
	"net/http"

	"context"

	"github.com/dhruvsaxena1998/aio/cmd/internal/database"
	"github.com/dhruvsaxena1998/aio/cmd/internal/helpers"
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
			helpers.ErrorResponse(w, "API Key is required", http.StatusUnauthorized)
			return
		}

		var user models.User
		err := database.DB.
			Preload("RoleGroup.Permissions").
			Where(&models.User{APIKey: apiKey, IsActive: true}).
			First(&user).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helpers.ErrorResponse(w, "Invalid API Key", http.StatusUnauthorized)
			} else {
				helpers.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
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
			helpers.ErrorResponse(w, "User not found in context", http.StatusUnauthorized)
			return
		}

		hasPermission, err := user.HasPermissions(database.DB, "allow:all")
		if err != nil || !hasPermission {
			helpers.ErrorResponse(w, "You do not have permission to access this resource", http.StatusForbidden)
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
