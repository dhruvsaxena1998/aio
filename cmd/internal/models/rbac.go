package models

type Permission struct {
	Base

	Name       string       `gorm:"not null;uniqueIndex " json:"name"`
	RoleGroups []*RoleGroup `gorm:"many2many:role_group_permissions;" json:"role_groups,omitempty"`
}

type RoleGroup struct {
	Base

	Name        string        `gorm:"not null" json:"name"`
	Description string        `json:"description"`
	Permissions []*Permission `gorm:"many2many:role_group_permissions;" json:"permissions,omitempty"`
	Users       []*User       `json:"users,omitempty"`
}
