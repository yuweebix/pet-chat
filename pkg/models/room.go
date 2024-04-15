package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name     string `gorm:"column:name;not null"`
	Users    []User `gorm:"many2many:user_rooms;"`
	Messages []Message
}

func (Room) TableName() string {
	return "rooms"
}
