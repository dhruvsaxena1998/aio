package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhruvsaxena1998/aio/cmd/internal/database"
	"github.com/dhruvsaxena1998/aio/cmd/internal/helpers"
	"github.com/dhruvsaxena1998/aio/cmd/internal/middlewares"
	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func CreateHabitHandler(w http.ResponseWriter, r *http.Request) {
	var input models.Habit
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helpers.ErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		helpers.ErrorResponse(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	input.UserID = userID

	if err := database.DB.Create(&input).Error; err != nil {
		helpers.ErrorResponse(w, "Failed to create habit", http.StatusInternalServerError)
		return
	}

	helpers.SuccessResponse(w, input, http.StatusCreated)
}

func GetHabitsHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		helpers.ErrorResponse(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var habits []models.Habit
	if err := database.DB.
		Where("user_id = ?", userID).
		Preload("Completions").
		Find(&habits).Error; err != nil {
		helpers.ErrorResponse(w, "Failed to fetch habits", http.StatusInternalServerError)
		return
	}

	helpers.SuccessResponse(w, habits, http.StatusOK)
}

func CreateCompletionHandler(w http.ResponseWriter, r *http.Request) {
	var input models.Completion
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helpers.ErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	habitIDStr := chi.URLParam(r, "id")
	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		helpers.ErrorResponse(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}
	input.HabitID = uint(habitID)

	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		helpers.ErrorResponse(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var habit models.Habit
	if err := database.DB.
		Where("id = ? AND user_id = ?", habitID, userID).
		First(&habit).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(w, "Habit not found or does not belong to the user", http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(w, "Failed to verify habit ownership", http.StatusInternalServerError)
		return
	}

	if err := database.DB.Create(&input).Error; err != nil {
		helpers.ErrorResponse(w, "Failed to create completion", http.StatusInternalServerError)
		return
	}

	helpers.SuccessResponse(w, input, http.StatusCreated)
}

func GetCompletionsHandler(w http.ResponseWriter, r *http.Request) {
	habitIDStr := chi.URLParam(r, "id")
	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		helpers.ErrorResponse(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}

	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		helpers.ErrorResponse(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	var habit models.Habit
	if err := database.DB.
		Where("id = ? AND user_id = ?", habitID, userID).
		First(&habit).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(w, "Habit not found or does not belong to the user", http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(w, "Failed to verify habit ownership", http.StatusInternalServerError)
		return
	}

	var completions []models.CompletionSerializer

	if err := database.DB.
		Model(&models.Completion{}).
		Select("id", "habit_id", "completed_at", "notes", "tags", "created_at").
		Where("habit_id = ?", habitID).
		Find(&completions).Error; err != nil {
		helpers.ErrorResponse(w, "Failed to fetch completions", http.StatusInternalServerError)
		return
	}

	helpers.SuccessResponse(w, completions, http.StatusOK)
}
