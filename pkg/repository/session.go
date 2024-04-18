package repository

import (
	"net/http"
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

func GetSessionByUser(db *gorm.DB, user *models.User) (*models.Session, error) {
	var session *models.Session
	result := db.Where("user_id = ?", user.ID).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func GetSessionByToken(db *gorm.DB, token string) (*models.Session, error) {
	var session *models.Session
	result := db.Where("token = ?", token).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func DeleteSession(db *gorm.DB, session *models.Session) error {
	result := db.Delete(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateSessionExpiry(db *gorm.DB, session *models.Session) error {
	t, err := strconv.Atoi(os.Getenv("EXPIRY"))
	if err != nil {
		return err
	}

	session.ExpiresAt = time.Now().Add(time.Minute * time.Duration(t))
	if err := db.Save(&session).Error; err != nil {
		return err
	}

	return nil
}

func GetAllSessions(db *gorm.DB) ([]models.Session, error) {
	var sessions []models.Session

	if err := db.Find(&sessions).Error; err != nil {
		return nil, err
	}

	return sessions, nil
}

func GetSessionByCookie(r *http.Request, db *gorm.DB) (*models.Session, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil, err
	}

	session, err := GetSessionByToken(db, cookie.Value)
	if err != nil {
		return nil, err
	}

	return session, nil
}
