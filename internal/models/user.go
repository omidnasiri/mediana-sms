package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	PasswordHash string
	SchoolId     uint
	School
	RoleId uint
	Role
}
