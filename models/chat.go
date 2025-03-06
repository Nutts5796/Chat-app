package models

import "github.com/jinzhu/gorm"

type Chat struct {
	gorm.Model
	Name         string `json:"name"`
	Type         string `json:"type"`
	Participants []User `gorm:"many2many:chat_users;"`
}
