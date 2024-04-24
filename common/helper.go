package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Hash the password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Generate JWT access token
func GenerateJWT(email, accessLevelID string) (string, error) {
	// Prepare the token with the claims
	var subject = struct {
		Email         string `json:"email"`
		AccessLevelID string `json:"access_level"`
	}{
		Email:         email,
		AccessLevelID: accessLevelID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Second * 24).Unix(),
	})

	return token.SignedString([]byte("secret"))
}

// Get access level id from JWT
func GetAccessLevelID(token *jwt.Token) string {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}

	sub, ok := claims["sub"].(map[string]interface{})
	if !ok {
		return ""
	}

	accessLevelID, ok := sub["access_level"].(string)
	if !ok {
		return ""
	}

	return accessLevelID
}
