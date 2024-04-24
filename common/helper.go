package common

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = os.Getenv("SECRET_KEY")

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

// Generate access token
func HashToken(email, permissionID string) (string, error) {
	// Prepare the token with the claims
	var subject = struct {
		Email        string `json:"email"`
		PermissionID string `json:"permission_id"`
	}{
		Email:        email,
		PermissionID: permissionID,
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})

	return tokenString.SignedString([]byte(SecretKey))
}

// Get permissionID from token
func GetPermissionID(tokenString *jwt.Token) string {
	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}

	sub, ok := claims["sub"].(map[string]interface{})
	if !ok {
		return ""
	}

	permissionID, ok := sub["permission_id"].(string)
	if !ok {
		return ""
	}

	return permissionID
}

// ParseToken decodes and validates a JWT token.
func ParseToken(tokenString string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", method)
		}
		return []byte(SecretKey), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, err
	}

	return parsedToken, nil
}
