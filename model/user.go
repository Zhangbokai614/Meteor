package model

import "time"

type APIUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	RID      uint   `json:"rid" binding:"required"`
}

type APIRole struct {
	Name            string `json:"name" binding:"required"`
	PerimssionsList []uint `json:"permissions_list" binding:"required"`
}

type User struct {
	ID        uint   `gorm:"primarykey"`
	RID       uint   `gorm:"not null"`
	Name      string `gorm:"unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	ID        uint   `gorm:"primarykey" json:"r_id"`
	Name      string `gorm:"unique" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RolePermissions struct {
	RID       uint `gorm:"primarykey"`
	PID       uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Permissions struct {
	ID                uint   `gorm:"primarykey"`
	RouterPermissions string `gorm:"unique"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
