package entity

import (
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
	TransactionTypeTopup    = "topup"
	TransactionTypeWithdraw = "withdraw"
	TransactionTypeTransfer = "transfer"
	TransactionTypeCheckBalance = "check_balance"
)

type Transaction struct {
	gorm.Model
	UserId            uuid.UUID         `gorm:"not null;type:uuid;unique"`
	WalletId          uuid.UUID         `gorm:"not null;type:uuid;unique"`
	Amount            float32           `gorm:"not null"`
	TransactionType   TransactionType            `gorm:"not null;type:varchar(255)"`
	TransactionStatus TransactionStatus `gorm:"not null;type:varchar(255)"`
}