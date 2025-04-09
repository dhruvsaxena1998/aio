package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhruvsaxena1998/aio/cmd/internal/database"
	"github.com/dhruvsaxena1998/aio/cmd/internal/middlewares"
	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
	"github.com/go-chi/chi/v5"
)

func CreateHabitHandler(w http.ResponseWriter, r *http.Request) {
	var input models.Habit
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	input.UserID = 1

	if err := database.DB.Create(&input).Error; err != nil {
		http.Error(w, "Failed to create habit", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}

func GetHabitsHandler(w http.ResponseWriter, r *http.Request) {
	var habits []models.Habit

	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	if err := database.DB.Where("user_id = ?", userID).Preload("Completions").Find(&habits).Error; err != nil {
		http.Error(w, "Failed to fetch habits", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(habits)
}

func CreateCompletionHandler(w http.ResponseWriter, r *http.Request) {
	var input models.Completion
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	habitID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}
	input.HabitID = uint(habitID)

	if err := database.DB.Create(&input).Error; err != nil {
		http.Error(w, "Failed to create completion", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}
