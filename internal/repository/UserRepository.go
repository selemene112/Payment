package repository

import (
	"fmt"
	"net/http"
	"payment/internal/entity"
	"payment/internal/excaption"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(data *entity.User, tx *gorm.DB, c *gin.Context) error {
	if err := tx.Create(&data).Error; err != nil {
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Create User Failed")
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func GetUserByEmail(data *entity.User, email string, tx *gorm.DB, c *gin.Context) error {
	if err := tx.Where("email = ?", email).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			excaption.ErrorHandler(c, http.StatusNotFound, err, "User Not Found")
			return fmt.Errorf("user not found: %w", err)
		}
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Get User Failed")
		return fmt.Errorf("failed to get user: %w", err)
	}
	return nil
}

func UpdateUserById(data *entity.User, ID uuid.UUID, tx *gorm.DB, c *gin.Context) error {
	if err := tx.Model(data).Where("id = ?", ID).Updates(&data).Error; err != nil {
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Update User Failed")
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
