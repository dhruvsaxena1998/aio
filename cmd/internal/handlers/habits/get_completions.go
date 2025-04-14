package habits

import (
	"net/http"
	"strconv"

	"github.com/dhruvsaxena1998/aio/cmd/internal/helpers"
	"github.com/dhruvsaxena1998/aio/cmd/internal/middlewares"
	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func (h *HabitsHandler) GetCompletionsHandler(w http.ResponseWriter, r *http.Request) {
	habitIDStr := chi.URLParam(r, "habitId")
	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		helpers.ErrorResponse(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}

	user, ok := middlewares.GetUserFromContext(r.Context())
	if !ok {
		helpers.ErrorResponse(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	var habit models.Habit
	if err := h.db.
		Where("id = ? AND user_id = ?", habitID, user.ID).
		First(&habit).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(w, "Habit not found or does not belong to the user", http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(w, "Failed to verify habit ownership", http.StatusInternalServerError)
		return
	}

	var completions []models.CompletionSerializer

	if err := h.db.
		Model(&models.Completion{}).
		Select("id", "habit_id", "completed_at", "notes", "tags", "created_at").
		Where("habit_id = ?", habitID).
		Find(&completions).Error; err != nil {
		helpers.ErrorResponse(w, "Failed to fetch completions", http.StatusInternalServerError)
		return
	}

	helpers.SuccessResponse(w, completions, http.StatusOK)
}
