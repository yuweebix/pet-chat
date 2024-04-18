package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Token     string `gorm:"column:token;unique;not null"`
	UserID    uint   `gorm:"column:user_id;not null"`
	User      User
	ExpiresAt time.Time `gorm:"column:expires_at"`
}

func (Session) TableName() string {
	return "sessions"
}

func (s Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}

func (s Session) IsAboutToExpire() bool {
	t := time.Now().Add(30 * time.Minute)
	return s.ExpiresAt.Before(t) || s.ExpiresAt.Equal(t)
}
