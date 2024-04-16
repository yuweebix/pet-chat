package repository

import (
	"errors"

	"github.com/yuweebix/pet-chat/pkg/models"
	"github.com/yuweebix/pet-chat/pkg/utils"
	"gorm.io/gorm"
)

var ErrInvalidPassword = errors.New("invalid password")

func CreateUser(db *gorm.DB, createdUser *models.UserCreate) error {
	hashedPassword, err := utils.HashPassword(createdUser.Password)
	if err != nil {
		return err
	}
	user := models.User{
		Email:          createdUser.Email,
		Username:       createdUser.Username,
		HashedPassword: hashedPassword,
		IsAdmin:        false,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(db *gorm.DB, loggedUser *models.UserLogin) (*models.User, error) {
	var user *models.User
	result := db.Where("username = ? OR email = ?", loggedUser.UsernameOrEmail, loggedUser.UsernameOrEmail).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if !utils.CheckPasswordHash(loggedUser.Password, user.HashedPassword) {
		return nil, ErrInvalidPassword
	}

	return user, nil
}

func GetUser(db *gorm.DB, username string) (*models.UserGet, error) {
	var user models.User
	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	user_info := &models.UserGet{
		Email:    user.Email,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		Rooms:    user.Rooms,
	}

	return user_info, nil
}
