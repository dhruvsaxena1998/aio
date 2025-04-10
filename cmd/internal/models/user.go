package models

type User struct {
	Base

	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	Habits   []Habit
}
