package habits

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhruvsaxena1998/aio/cmd/internal/helpers"
	"github.com/dhruvsaxena1998/aio/cmd/internal/middlewares"
	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func (h *HabitsHandler) CreateCompletionHandler(w http.ResponseWriter, r *http.Request) {
	var input models.Completion
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helpers.ErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	habitIDStr := chi.URLParam(r, "habitId")
	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		helpers.ErrorResponse(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}
	input.HabitID = uint(habitID)

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

	if err := h.db.Create(&input).Error; err != nil {
		helpers.ErrorResponse(w, "Failed to create completion", http.StatusInternalServerError)
		return
	}

	helpers.SuccessResponse(w, input, http.StatusCreated)
}
