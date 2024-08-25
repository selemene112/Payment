package csvupload

import (
	"fmt"
	"payment/internal/entity"
	paymentmodel "payment/internal/model/paymentModel"
	"strconv"

	"github.com/google/uuid"
)

func MapCSVRecordToPaymentRequest(record []string) (paymentmodel.PaymentRequest, error) {
	amount, err := strconv.ParseFloat(record[0], 32)
	if err != nil {
		return paymentmodel.PaymentRequest{}, fmt.Errorf("invalid amount format: %v", err)
	}
	var transactionType entity.TransactionType
	transactionType, err = stringToTransactionType(record[1])
	if err != nil {
		return paymentmodel.PaymentRequest{}, fmt.Errorf("invalid transaction type format: %v", err)
	}
	var destinationWalletId uuid.UUID

	if record[2] != "" {
		destinationWalletId, err = uuid.Parse(record[2])
		if err != nil {
			return paymentmodel.PaymentRequest{}, fmt.Errorf("invalid destination wallet id format: %v", err)
		}
	}

	request := paymentmodel.PaymentRequest{
		Amount:              float32(amount),
		TransactionType:     transactionType,
		DestinationWalletId: destinationWalletId,
	}

	return request, nil
}

func stringToTransactionType(transactionType string) (entity.TransactionType, error) {
	switch transactionType {
	case "topup":
		return entity.TransactionTypeTopup, nil
	case "transfer":
		return entity.TransactionTypeTransfer, nil
	case "withdraw":
		return entity.TransactionTypeWithdraw, nil
	case "check_balance":
		return entity.TransactionTypeCheckBalance, nil
	default:
		return "", fmt.Errorf("invalid transaction type: %s", transactionType)
	}
}
