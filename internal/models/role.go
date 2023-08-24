package models

import "gorm.io/gorm"

type RoleType = string

const (
	ROLE_ADMIN      RoleType = "admin"
	ROLE_HEADMASTER RoleType = "headmaster"
	ROLE_TEACHER    RoleType = "teacher"
	ROLE_STUDENT    RoleType = "student"
)

type Role struct {
	gorm.Model
	Title RoleType `json:"title"`
	Users []*User  `json:"users,omitempty"`
}
