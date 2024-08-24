package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model	
	Amount float32 `gorm:"not null"`
	UserId uuid.UUID `gorm:"not null;type:uuid;unique"`
	Transactions []Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}