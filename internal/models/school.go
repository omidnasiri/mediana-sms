package models

import "gorm.io/gorm"

type School struct {
	gorm.Model
	Title        string  `json:"title"`
	HeadmasterId uint    `json:"headmaster_id"`
	Users        []*User `json:"users,omitempty"`
}
