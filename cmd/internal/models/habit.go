package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type HabitType int

const (
	GoodHabit HabitType = 0
	BadHabit  HabitType = 1
)

func (ht HabitType) IsValid() bool {
	return ht == GoodHabit || ht == BadHabit
}

type Habit struct {
	gorm.Model

	UserID      uint      `gorm:"not null;index" json:"user_id"`
	Name        string    `gorm:"not null" json:"name"`
	Type        HabitType `gorm:"not null" json:"type"`
	IsActive    bool      `gorm:"not null;default:true" json:"is_active"`
	Completions []Completion
}

// BeforeCreate hook to validate HabitType before saving to the database
func (h *Habit) BeforeCreate(tx *gorm.DB) error {
	if !h.Type.IsValid() {
		return fmt.Errorf("invalid habit type: %d", h.Type)
	}
	return nil
}

// BeforeUpdate hook to validate HabitType before updating the database
func (h *Habit) BeforeUpdate(tx *gorm.DB) error {
	if !h.Type.IsValid() {
		return fmt.Errorf("invalid habit type: %d", h.Type)
	}
	return nil
}

type Completion struct {
	gorm.Model

	HabitID     uint      `gorm:"not null;index" json:"habit_id"`
	CompletedAt time.Time `gorm:"not null;index;default:CURRENT_TIMESTAMP" json:"completed_at"`
	Notes       string    `gorm:"type:text" json:"notes"`
	Tags        string    `gorm:"type:text" json:"tags"` // Comma-separated tags
}
