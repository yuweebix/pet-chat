package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"column:email;unique;not null"`
	Username       string `gorm:"column:username;unique;not null"`
	HashedPassword string `gorm:"column:password;unique;not null"`
	IsAdmin        bool   `gorm:"column:is_admin;not null"`
	Rooms          []Room `gorm:"many2many:user_rooms;"`
	Messages       []Message
}

func (User) TableName() string {
	return "users"
}
