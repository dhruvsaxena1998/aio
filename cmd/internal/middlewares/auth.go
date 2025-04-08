package middlewares

import "net/http"

const ADMIN_API_KEY = "H6WFaOZ7-8JWFgXo-P10UNPzG-NIKQs3DU-WU4SB6D1"

func OnlyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey != ADMIN_API_KEY {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
