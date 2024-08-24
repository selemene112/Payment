package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
	Photo string `gorm:"type:varchar(255)"`
	Wallet Wallet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transactions []Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}