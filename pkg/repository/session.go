package repository

import (
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/yuweebix/pet-chat/pkg/models"
	"gorm.io/gorm"
)

func CreateSession(db *gorm.DB, user *models.User) error {
	token := uuid.NewString()
	expiryTime, _ := strconv.Atoi(os.Getenv("EXPIRY"))
	expiry := time.Now().Add(time.Minute * time.Duration(expiryTime))

	session := &models.Session{
		Token:     token,
		User:      *user,
		ExpiresAt: expiry,
	}

	if err := db.Create(&session).Error; err != nil {
		return err
	}
	return nil
}

func GetSession(db *gorm.DB, user *models.User) (*models.Session, error) {
	var session *models.Session
	result := db.Where("user_id = ?", user.ID).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}
