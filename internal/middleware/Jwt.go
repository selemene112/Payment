package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func GenerateToken(id uuid.UUID, name string, email string) string {
	expirationTime := time.Now().Add(5 * time.Minute).Unix()
	claims := jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   expirationTime,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := parseToken.SignedString([]byte("secret"))
	return token
}