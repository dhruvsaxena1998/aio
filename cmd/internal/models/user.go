package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	Habits   []Habit
}
