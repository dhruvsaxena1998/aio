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

func (u *User) loadPermissions(db *gorm.DB) error {
	if u.RoleGroupID == 0 {
		return fmt.Errorf("user has no role group assigned")
	}

	if u.RoleGroup.ID == 0 || u.RoleGroup.Permissions == nil {
		if err := db.Preload("RoleGroup.Permissions").First(u, u.ID).Error; err != nil {
			return fmt.Errorf("failed to load user with role group and permissions: %w", err)
		}
	}

	if len(u.RoleGroup.Permissions) == 0 {
		return fmt.Errorf("role group '%s' has no permissions assigned", u.RoleGroup.Name)
	}

	return nil
}

func (u *User) HasPermissions(db *gorm.DB, permissionName string) (bool, error) {
	if err := u.loadPermissions(db); err != nil {
		return false, err
	}

	for _, perm := range u.RoleGroup.Permissions {
		if perm.Name == permissionName {
			return true, nil
		}
	}

	return false, nil
}

func (u *User) GetPermissions(db *gorm.DB) ([]string, error) {
	if err := u.loadPermissions(db); err != nil {
		return nil, err
	}

	permissions := make([]string, len(u.RoleGroup.Permissions))
	for i, perm := range u.RoleGroup.Permissions {
		permissions[i] = perm.Name
	}

	return permissions, nil
}
