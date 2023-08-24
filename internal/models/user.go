package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string
	SchoolId     uint
	School       *School
	RoleId       uint
	Role         *Role
}
