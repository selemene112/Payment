package helper

import (
	"net/http"
	"payment/internal/excaption"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateInput(input interface{}, c *gin.Context) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		excaption.ErrorHandler(c, http.StatusBadRequest, err, "Input Data Not Valid")
		return err
	}
	return nil
}

func ValidateInputCustom(input interface{}, c *gin.Context, tagName string, function func(fl validator.FieldLevel) bool) error {
	validate := validator.New()
	validate.RegisterValidation(tagName, function)
	err := validate.Struct(input)
	if err != nil {
		excaption.ErrorHandler(c, http.StatusBadRequest, err, "Input Data Not Valid")
		return err
	}
	return nil
}