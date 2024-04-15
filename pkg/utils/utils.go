package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func CheckUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,}$`)
	return re.MatchString(username)
}

func CheckPassword(password string) bool {
	re := regexp.MustCompile(`^.{8,}$`)
	return re.MatchString(password)
}
