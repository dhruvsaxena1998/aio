package habits

import "gorm.io/gorm"

type HabitsHandler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *HabitsHandler {
	return &HabitsHandler{
		db: db,
	}
}
