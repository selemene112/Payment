package controller

import (
	"net/http"
	"payment/internal/excaption"
	"payment/internal/model"
	"payment/internal/model/authModel"
	authservice "payment/internal/service/authService"

	"github.com/gin-gonic/gin"
)

func RegisterUserController(c *gin.Context) {
	var request authModel.AuthRegisterUserRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		excaption.ErrorHandler(c, http.StatusBadRequest, err, "Input Data Not Valid")
		return
	}

	response , err := authservice.RegisterUserService(request, c); if err != nil {
		return
	}

	c.JSON(http.StatusCreated,model.StandarResponse{
		Status:  http.StatusCreated,
		Error:   "",
		Message: "User Created",
		Result:  response,
	})
}

func LoginUserController(c *gin.Context) {
	var request authModel.AuthLoginUserRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		excaption.ErrorHandler(c, http.StatusBadRequest, err, "Input Data Not Valid")
		return
	}

	response , err := authservice.LoginUserService(request, c); if err != nil {
		return
	}

	c.JSON(http.StatusOK,model.StandarResponse{
		Status:  http.StatusOK,
		Error:   "",
		Message: "Login Success",
		Result:  response,
	})
}