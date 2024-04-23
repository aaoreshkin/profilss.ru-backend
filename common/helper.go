package common

import (
	"golang.org/x/crypto/bcrypt"
)

type Payload struct {
	Email         string `json:"email"`
	AccessLevelID string `json:"access_level_id"`
}

// Hash the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
