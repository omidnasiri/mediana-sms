package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	UserId   uint
	User     *User
	Teachers []*Teacher `gorm:"many2many:teacher_student" json:"teachers,omitempty"`
}
