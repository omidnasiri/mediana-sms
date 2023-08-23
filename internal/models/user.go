package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	PasswordHash string
	SchoolId     *uint
	School       *School
	RoleId       uint
	Role         *Role
}
