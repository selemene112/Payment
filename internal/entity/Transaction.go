package entity

import (
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionStatus string

const (
	TransactionStatusSuccess   = "success"
	TransactionStatusFailed    = "failed"
	TransactionStatusCancelled = "cancelled"
)

type TransactionType string

const (
	TransactionTypeTopup        = "topup"
	TransactionTypeWithdraw     = "withdraw"
	TransactionTypeTransfer     = "transfer"
	TransactionTypeCheckBalance = "check_balance"
)

type Transaction struct {
	gorm.Model
	UserId              uuid.UUID         `gorm:"not null;type:uuid;"`
	WalletId            uuid.UUID         `gorm:"not null;type:uuid;"`
	DestinationWalletId uuid.UUID         `json:"destination_wallet_id" binding:"required"`
	Request             json.RawMessage   `gorm:"not null"`
	Response            json.RawMessage   `gorm:"not null"`
	Amount              float32           `gorm:"not null"`
	TransactionType     TransactionType   `gorm:"not null;type:varchar(255)"`
	TransactionStatus   TransactionStatus `gorm:"not null;type:varchar(255)"`
}
