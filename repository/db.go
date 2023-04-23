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

	// config := models.Config{
	// 	DB_Username: "root",
	// 	DB_Password: "polisi21",
	// 	DB_Port:     "3306",
	// 	DB_Host:     "localhost",
	// 	DB_Name:     "mini_project",
	// }

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	config.DB_Username,
	// 	config.DB_Password,
	// 	config.DB_Host,
	// 	config.DB_Port,
	// 	config.DB_Name,
	// )

	// var err error
	// DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
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
