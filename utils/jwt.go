// utils/jwt.go
package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a JWT token
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte("secret"))
}
