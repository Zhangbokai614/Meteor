package model

import "time"

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type APIUser struct {
	Name                       string `json:"name" binding:"required"`
	Password                   string `json:"password" binding:"required"`
	Unit                       *int   `json:"unit" binding:"required"`
	StageEarlyPermissions      *int   `json:"early_permissions" binding:"required"`
	StageManagementPermissions *int   `json:"management_permissions" binding:"required"`
	StageAccountsPermissions   *int   `json:"accounts_permissions" binding:"required"`
}
