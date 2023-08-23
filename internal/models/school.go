package models

import "gorm.io/gorm"

type School struct {
	gorm.Model
	Title    string
	Admin    *User
	Teachers []*Teacher
	Student  []*Student
}
