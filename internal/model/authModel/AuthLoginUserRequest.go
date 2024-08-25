package authModel

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type AuthLoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password" validate:"required,ValidatePassword"`
}

const (
	minPasswordLength = 8
	maxPasswordLength = 32
)

func PasswordValidator(fl validator.FieldLevel) bool {
	var passwordRegex = regexp.MustCompile(`[A-Z]`).MatchString(fl.Field().String()) && regexp.MustCompile(`[a-z]`).MatchString(fl.Field().String()) && regexp.MustCompile(`[0-9]`).MatchString(fl.Field().String()) && regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':"\\|,.<>\/?~]`).MatchString(fl.Field().String())
	return len(fl.Field().String()) >= minPasswordLength && len(fl.Field().String()) <= maxPasswordLength && passwordRegex
}
