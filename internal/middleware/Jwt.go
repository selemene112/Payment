package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"payment/internal/excaption"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func GenerateToken(id uuid.UUID, name string, email string) string {
	env := godotenv.Load()
	if env != nil {
		panic("Failed to load .env file")
	}
	var secretkey = os.Getenv("SECRET_KEY")
	expirationTime := time.Now().Add(100 * time.Minute).Unix()
	claims := jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   expirationTime,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := parseToken.SignedString([]byte(secretkey))
	return token
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	env := godotenv.Load()
	if env != nil {
		panic("Failed to load .env file")
	}
	var secretkey = os.Getenv("SECRET_KEY")
	errResponse := errors.New("Input Token First")
	headerToken := c.Request.Header.Get("authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errResponse
	}
	stringToken := strings.Split(headerToken, " ")[1]
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token expired")
			}
		}
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errResponse
	}
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, errors.New("token expired")
		}
	}

	return claims, nil
}

func VerifyAndNext(c *gin.Context) {
	claims, err := VerifyToken(c)
	if err != nil {
		fmt.Println("Token verifikasi gagal:", err)
		excaption.ErrorHandler(c, http.StatusUnauthorized, err, "Unauthorized")
		c.Abort()
		return
	}

	if claims == nil {
		fmt.Println("Claims kosong di middleware")
		excaption.ErrorHandler(c, http.StatusUnauthorized, errors.New("claims kosong"), "Unauthorized")
		c.Abort()
		return
	}

	fmt.Println("Yes Lolos Midleware")
	c.Set("claims", claims)
	c.Next()
}
