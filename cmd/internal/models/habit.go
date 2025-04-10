package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type HabitType int

const (
	GoodHabit HabitType = iota
	BadHabit
)

func (ht HabitType) String() string {
	switch ht {
	case GoodHabit:
		return "good"
	case BadHabit:
		return "bad"
	default:
		return fmt.Sprintf("unknown(%d)", ht)
	}
}

func (ht HabitType) IsValid() bool {
	return ht == GoodHabit || ht == BadHabit
}

type Habit struct {
	Base

	UserID      uint      `gorm:"not null;index" json:"user_id"`
	Name        string    `gorm:"not null" json:"name"`
	Type        HabitType `gorm:"not null" json:"type"`
	IsActive    bool      `gorm:"not null;default:true" json:"is_active"`
	Completions []Completion
}

func (h *Habit) validateHabitType() error {
	if !h.Type.IsValid() {
		return fmt.Errorf("invalid habit type: %d", h.Type)
	}
	return nil
}

func (h *Habit) BeforeCreate(tx *gorm.DB) error {
	return h.validateHabitType()
}

func (h *Habit) BeforeUpdate(tx *gorm.DB) error {
	return h.validateHabitType()
}

type Completion struct {
	Base

	HabitID     uint      `gorm:"not null;index" json:"habit_id"`
	CompletedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"completed_at"`
	Notes       string    `gorm:"type:text" json:"notes"`
	Tags        string    `gorm:"type:text" json:"tags"`
}

type CompletionSerializer struct {
	ID          uint      `json:"id"`
	HabitID     uint      `json:"habit_id"`
	CompletedAt time.Time `json:"completed_at"`
	Notes       string    `json:"notes"`
	Tags        string    `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
}
