package habits

import (
	"net/http"

	"github.com/dhruvsaxena1998/aio/cmd/internal/helpers"
	"github.com/dhruvsaxena1998/aio/cmd/internal/middlewares"
	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
)

func (h *HabitsHandler) GetHabitsHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := middlewares.GetUserFromContext(r.Context())
	if !ok {
		helpers.ErrorResponse(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var habits []models.Habit
	if err := h.db.
		Where("user_id = ?", user.ID).
		Preload("Completions").
		Find(&habits).Error; err != nil {
		helpers.ErrorResponse(w, "Failed to fetch habits", http.StatusInternalServerError)
		return
	}

	helpers.SuccessResponse(w, habits, http.StatusOK)
}
