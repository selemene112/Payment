package authservice

import (
	"fmt"
	"payment/internal/entity"
	"payment/internal/helper"
	"payment/internal/middleware"
	"payment/internal/model/authModel"
	"payment/internal/repository"
	"payment/pkg/config"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUserService(request authModel.AuthLoginUserRequest, c *gin.Context) (authModel.TokenResponse, error) {
	err := helper.ValidateInputCustom(request, c, "ValidatePassword", authModel.PasswordValidator)
	if err != nil {
		return authModel.TokenResponse{}, err
	}
	var token string
	var User entity.User
	err = config.GetPgsqlDB().Transaction(func(tx *gorm.DB) error {

		if err := repository.GetUserByEmail(&User, request.Email, tx, c); err != nil {
			return err
		}

		var wg sync.WaitGroup
		errChan := make(chan error, 2+1)
		wg.Add(1)
		go func () {
			defer wg.Done()
			if !middleware.CheckPasswordHash([]byte(User.Password), []byte(request.Password)) {
				errChan <- fmt.Errorf("password not match")
				return 
			}
			errChan <- nil
			return		
		}()

		wg.Add(1)
		go func () {
			defer wg.Done()
			token = middleware.GenerateToken(User.ID, User.Name, User.Email)
			errChan <- nil
			return
		}()

		wg.Wait()
		close(errChan)
		for err := range errChan {
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return authModel.TokenResponse{}, err

	}


	return authModel.TokenResponse{
		AccessToken:  token,
	}, nil

}

func RegisterUserService(request authModel.AuthRegisterUserRequest, c *gin.Context)(authModel.AuthRegisterUserResponse, error) {
	err := helper.ValidateInputCustom(request, c, "ValidatePassword", authModel.PasswordValidator)
	if err != nil {
		return authModel.AuthRegisterUserResponse{}, err
	}

	User := entity.User{}
	err = config.GetPgsqlDB().Transaction(func(tx *gorm.DB) error {
		passwordHash := middleware.Bcrypt(request.Password)
		User = entity.User{
			Name: request.Name,
			Email: request.Email,
			Password: passwordHash,
		}
		err := repository.CreateUser(&User, tx, c); if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return authModel.AuthRegisterUserResponse{}, err
	}
	return authModel.AuthRegisterUserResponse{
		Name: User.Name,
		Email: User.Email,
	}, nil

}