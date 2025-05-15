package handler

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(raw_password string) (string, error) {
	hashed_bytes, err := bcrypt.GenerateFromPassword([]byte(raw_password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed_bytes), err
}

func IsCorrectPassword(hashed_password, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	return err == nil
}