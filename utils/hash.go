package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func DecodePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))

	return err == nil
}
