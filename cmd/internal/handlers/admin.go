package handlers

import (
	"net/http"

	"github.com/dhruvsaxena1998/aio/cmd/internal/helpers"
)

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SuccessResponse(w, map[string]string{
		"message": "You are an admin 🎉",
	}, http.StatusOK)
}
