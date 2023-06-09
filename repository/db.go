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
	dsn := "host=localhost user=postgres password= dbname=mini_project port=5431 sslmode=disable TimeZone=Asia/Jakarta"
	// dsn := "host=docker.for.mac.localhost user=postgres password= dbname=mini_project port=5431 sslmode=disable TimeZone=Asia/Jakarta"
	// dsn := "host=db-mini-project.cywocig1ynmq.ap-southeast-2.rds.amazonaws.com user=postgres password=postgres-password dbname=mini_project port=5432 sslmode=disable TimeZone=Asia/Jakarta"
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
	DB.AutoMigrate(&models.ShoesSize{})
	DB.AutoMigrate(&models.Carts{})
	DB.AutoMigrate(&models.PaymentMethod{})
	DB.AutoMigrate(&models.Shipping{})
	DB.AutoMigrate(&models.Transaction{})
	DB.AutoMigrate(&models.TransactionDetail{})
}
