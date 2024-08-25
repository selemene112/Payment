package repository

import (
	"fmt"
	"net/http"
	"payment/internal/entity"
	"payment/internal/excaption"
	"payment/pkg/config"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(data *entity.Transaction, c *gin.Context) error {
	if err := config.GetPgsqlDB().Create(&data).Error; err != nil {
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Create Transaction Failed")
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	return nil
}