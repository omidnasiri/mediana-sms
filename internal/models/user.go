package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string
	SchoolId     *uint   `json:"school_id,omitempty"`
	School       *School `json:"school,omitempty"`
	RoleId       uint    `json:"role_id,omitempty"`
	Role         *Role   `json:"role,omitempty"`
}
