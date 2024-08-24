package excaption

import (
	"payment/internal/model"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, status int, err error, message string) {
	c.JSON(status, model.StandarResponse{
		Status:  status,
		Error:   err.Error(),
		Message: message,
		Result:  nil,
	})
}