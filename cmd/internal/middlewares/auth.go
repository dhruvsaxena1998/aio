package middlewares

import (
	"net/http"
	"os"

	"context"
)

type ContextKey string

const (
	UserIdContextKey ContextKey = "user_id"
)

func OnlyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey != os.Getenv("API_KEY") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID := uint(1)

		ctx := context.WithValue(r.Context(), UserIdContextKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserIDFromContext(ctx context.Context) (uint, bool) {
	if userID, ok := ctx.Value(UserIdContextKey).(uint); ok {
		return userID, true
	}
	return 0, false
}
