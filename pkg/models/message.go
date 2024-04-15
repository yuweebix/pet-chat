package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content string `gorm:"column:content;not null"`
	UserID  uint   `gorm:"column:user_id;not null"`
	RoomID  uint   `gorm:"column:room_id;not null"`
	User    User
	Room    Room
}

func (Message) TableName() string {
	return "messages"
}
