package habits

import (
	"encoding/json"
	"net/http"

	"github.com/dhruvsaxena1998/aio/cmd/internal/helpers"
	"github.com/dhruvsaxena1998/aio/cmd/internal/middlewares"
	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
)

func (h *HabitsHandler) CreateHabitHandler(w http.ResponseWriter, r *http.Request) {
	var input models.HabitRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helpers.ErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, ok := middlewares.GetUserFromContext(r.Context())
	if !ok {
		helpers.ErrorResponse(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	habitType, err := models.StringToHabitType(input.Type)
	if err != nil {
		helpers.ErrorResponse(w, "Invalid habit type", http.StatusBadRequest)
		return
	}

	payload := models.Habit{
		UserID: user.ID,
		Name:   input.Name,
		Type:   habitType,
	}

	if err := h.db.Create(&payload).Error; err != nil {
		helpers.ErrorResponse(w, "Failed to create habit", http.StatusInternalServerError)
		return
	}

	helpers.SuccessResponse(w, input, http.StatusCreated)
}
