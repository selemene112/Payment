package helper

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Token struct {
	Id    uuid.UUID
	Nama  string
	Email string
}

func GetToken(c *gin.Context) (Token, error) {
	token, exists := c.Get("claims")
	if !exists {
		return Token{}, errors.New("Sign in to Proceed")
	}

	idConvert := token.(jwt.MapClaims)
	idStr := idConvert["id"].(string)
	id, _ := uuid.Parse(idStr)
	return Token{
			Id:    id,
			Email: idConvert["email"].(string),
			Nama:  idConvert["name"].(string),
		},
		nil
}