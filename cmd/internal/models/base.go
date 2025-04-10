package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
