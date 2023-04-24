package repository

import (
	"main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	dsn := "host=localhost user=postgres password= dbname=mini_project port=5431 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Shoes{})
	DB.AutoMigrate(&models.ShoesDetail{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.CartItems{})
	DB.AutoMigrate(&models.PaymentMethod{})
	DB.AutoMigrate(&models.Shipping{})
	DB.AutoMigrate(&models.Transaction{})
	DB.AutoMigrate(&models.TransactionDetail{})
}
