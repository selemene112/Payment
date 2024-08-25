package config

import (
	"fmt"
	"payment/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "lele123"
	DB_NAME = "payment"
	db      *gorm.DB
	err     error
)

func ConnectPgsql() {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&entity.User{}, &entity.Wallet{}, &entity.Transaction{})
	fmt.Println("Database Migrated")

}

func GetPgsqlDB() *gorm.DB {
	return db
}
