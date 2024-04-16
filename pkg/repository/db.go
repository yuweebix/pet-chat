package repository

import (
	"fmt"
	"os"

	"github.com/yuweebix/pet-chat/pkg/models"
	"github.com/yuweebix/pet-chat/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Message{}, &models.Room{}, &models.Session{}, &models.User{}); err != nil {
		return nil, err
	}

	if err := initAdmin(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initAdmin(db *gorm.DB) error {
	email, username, password := os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_PASSWORD")
	hashedPassword, _ := utils.HashPassword(password)

	var user models.User
	if err := db.Where("email = ? AND username = ?", email, username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			admin := models.User{
				Email:          email,
				Username:       username,
				HashedPassword: hashedPassword,
				IsAdmin:        true,
			}

			tx := db.Begin()
			result := tx.Save(&admin)
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			} else {
				tx.Commit()
			}
		} else {
			return err
		}
	}
	return nil
}
