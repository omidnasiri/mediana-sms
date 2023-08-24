package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	UserId   uint
	User     *User
	Students []*Student `gorm:"many2many:teacher_student" json:"students,omitempty"`
}
