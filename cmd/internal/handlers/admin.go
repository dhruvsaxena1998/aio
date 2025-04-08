package handlers

import (
	"encoding/json"
	"net/http"
)

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "You are an admin 🎉",
	})
}
