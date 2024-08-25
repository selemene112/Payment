package controller

import (
	"encoding/csv"
	"io"
	"net/http"
	"payment/internal/excaption"
	"payment/internal/model"
	paymentmodel "payment/internal/model/paymentModel"
	csvupload "payment/internal/service/csvUpload"
	paymentservice "payment/internal/service/paymentService"

	"github.com/gin-gonic/gin"
)

func ProcessCSVHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		excaption.ErrorHandler(c, http.StatusBadRequest, err, "Failed to read file")
		return
	}
	defer file.Close()

	// Buat reader CSV
	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Failed to read CSV header")
		return
	}
	var paymentRequests []paymentmodel.PaymentRequest
	for {
		// Membaca baris CSV
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			excaption.ErrorHandler(c, http.StatusBadRequest, err, "Failed to read CSV record")
			return
		}
		// Memproses baris CSV apabila tidak kosong
		if len(record) == 0 || record[0] == "" {
			continue
		}
		request, err := csvupload.MapCSVRecordToPaymentRequest(record)
		if err != nil {
			excaption.ErrorHandler(c, http.StatusBadRequest, err, "Invalid CSV record")
			return
		}
		// kumpulkan hasil CSV ke dalam array
		paymentRequests = append(paymentRequests, request)
	}
	responses, err := paymentservice.MultiPaymentService(paymentRequests, c)
	if err != nil {
		excaption.ErrorHandler(c, http.StatusInternalServerError, err, "Failed to process CSV")
		return
	}
	c.JSON(http.StatusOK, model.StandarResponse{
		Status:  http.StatusOK,
		Error:   "",
		Message: "Payment Success",
		Result:  responses,
	})
}
