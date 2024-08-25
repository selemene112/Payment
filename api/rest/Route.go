package rest

import (
	"payment/internal/controller"
	"payment/internal/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	defaultRoute := r.Group("/api/v1.0.0")
	{

		AuthUsersRoute := defaultRoute.Group("/auth/users")
		{

			AuthUsersRoute.POST("/register", controller.RegisterUserController)
			AuthUsersRoute.POST("/login", controller.LoginUserController)
		}

		PaymentsRoute := defaultRoute.Group("/payments")
		{
			PaymentsRoute.POST("/payment", middleware.VerifyAndNext, controller.PaymentgController)
		}
	}

	return r
}
