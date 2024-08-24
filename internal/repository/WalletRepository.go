package repository

import (
	"fmt"
	"net/http"
	"payment/internal/entity"
	"payment/internal/excaption"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateWallet(data *entity.Wallet, tx *gorm.DB, c *gin.Context) error {
	if err := tx.Create(&data).Error; err != nil {
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Create Wallet Failed")
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	return nil
}