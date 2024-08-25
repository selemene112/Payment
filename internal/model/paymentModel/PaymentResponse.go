package paymentmodel

import (
	"payment/internal/entity"
	"time"

	"github.com/google/uuid"
)

type PaymentResponse struct {
	UserId          uuid.UUID  `json:"user_id"`
	WalletId        uuid.UUID  `json:"wallet_id"`
	TransactionDate time.Time  `json:"transaction_date"`
	BalanceBefore    float32    `json:"balance_before"`
	Amount          float32 `json:"amount"`
	BalanceAfter     float32    `json:"balance_after"`
	TransactionType entity.TransactionType  `json:"transaction_type"`
	DestinationWalletId        uuid.UUID `json:"destination_wallet_id"`
}