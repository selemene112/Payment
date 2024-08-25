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

func CreateWallet(data *entity.Wallet, tx *gorm.DB, c *gin.Context) error {
	if err := tx.Create(&data).Error; err != nil {
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Create Wallet Failed")
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	return nil
}

func GetWalletByUserId(data *entity.Wallet, userId uuid.UUID, tx *gorm.DB, c *gin.Context) error {
	if err := tx.Where("user_id = ?", userId).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			excaption.ErrorHandler(c, http.StatusNotFound, err, "Wallet Not Found")
			return fmt.Errorf("wallet not found: %w", err)
		}
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Get Wallet Failed")
		return fmt.Errorf("failed to get wallet: %w", err)
	}
	return nil
}

func GetWalletById(data *entity.Wallet, id uuid.UUID, tx *gorm.DB, c *gin.Context) error {
	if err := tx.Where("id = ?", id).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			excaption.ErrorHandler(c, http.StatusNotFound, err, "invalid destination wallet")
			return fmt.Errorf("invalid destination wallet: %w", err)
		}
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Get Wallet Failed")
		return fmt.Errorf("failed to get wallet: %w", err)
	}
	return nil
}

func UpdateWalletAmount(wallet *entity.Wallet, ID uuid.UUID, amount float32, tx *gorm.DB, c *gin.Context) error {
	if err := tx.Model(wallet).Where("id = ?", ID).Update("amount", amount).Error; err != nil {
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Update Wallet Failed")
		return fmt.Errorf("failed to update wallet: %w", err)
	}
	return nil
}
