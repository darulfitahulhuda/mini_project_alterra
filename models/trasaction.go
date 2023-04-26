package models

import (
	"main/dto"

	"gorm.io/gorm"
)

const PAYMENT_STATUS_WAITING = "waiting"

// const PAYMENT_STATUS_SUCCESS = "success"
const TRANSACTION_WAITING_PAYMENT = "waiting payment"

type Transaction struct {
	*gorm.Model
	UserId            uint                `json:"user_id" form:"user_id" gorm:"column:user_id"`
	TotalPrice        float64             `json:"total_price" form:"total_price" gorm:"column:total_price"`
	Status            string              `json:"status" form:"status"`
	PaymentMethod     PaymentMethod       `json:"payment_method" form:"payment_method" gorm:"foreignkey:transaction_id"`
	TransactionDetail []TransactionDetail `json:"detail" form:"detail" gorm:"foreignkey:trasaction_id"`
	Shipping          Shipping            `json:"shipping" form:"shipping"`
	DeletedAt         gorm.DeletedAt      `gorm:"index"`
}

type TransactionDetail struct {
	*gorm.Model
	TransactionId uint           `json:"trasaction_id" form:"trasaction_id" gorm:"column:trasaction_id"`
	ShoesId       uint           `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Qty           int            `json:"qty" form:"qty"`
	Price         float64        `json:"price" form:"price"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type TransactionResponse struct {
	Message string          `json:"message"`
	Status  int             `json:"status"`
	Data    dto.Transaction `json:"data"`
}
