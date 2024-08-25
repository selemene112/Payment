package controller

import (
	"net/http"
	"payment/internal/excaption"
	"payment/internal/model"
	paymentmodel "payment/internal/model/paymentModel"
	paymentservice "payment/internal/service/paymentService"

	"github.com/gin-gonic/gin"
)

func PaymentgController(c *gin.Context) {
	request := paymentmodel.PaymentRequest{}
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		excaption.ErrorHandler(c, http.StatusBadRequest, err, "Input Data Not Valid")
		return
	}

	response, err := paymentservice.PaymentTransaction(request, c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, model.StandarResponse{
		Status:  http.StatusOK,
		Error:   "",
		Message: "Payment Success",
		Result:  response,
	})

}
