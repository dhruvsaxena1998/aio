package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Base

	Name        string    `json:"name"`
	Email       string    `gorm:"uniqueIndex" json:"email"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	APIKey      string    `gorm:"uniqueIndex" json:"api_key"`
	RoleGroupID uint      `json:"role_group_id"`
	RoleGroup   RoleGroup `gorm:"foreignKey:RoleGroupID" json:"role_group"`
	Habits      []Habit
}

func (u *User) HasPermissions(db *gorm.DB, permsission string) (bool, error) {
	var roleGroup RoleGroup
	if err := db.Preload("Permissions").First(&roleGroup, u.RoleGroupID).Error; err != nil {
		return false, fmt.Errorf("failed to get role group: %w", err)
	}

	for _, perm := range roleGroup.Permissions {
		if perm != nil && perm.Name == permsission {
			return true, nil
		}
	}

	return false, nil
}
