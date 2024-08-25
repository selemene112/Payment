package paymentmodel

import (
	"payment/internal/entity"

	"github.com/google/uuid"
)

type PaymentRequest struct {
	Amount              float32                `json:"amount" binding:"required"`
	TransactionType     entity.TransactionType `json:"transaction_type" binding:"required"`
	DestinationWalletId uuid.UUID              `json:"destination_wallet_id"`
}
