package models

import (
	"gorm.io/gorm"
)

const PAYMENT_STATUS_WAITING = "Waiting"
const PAYMENT_STATUS_SUCCESS = "Success"

const TRANSACTION_PAYMENT_WAITING = "Payment waiting"
const TRANSACTION_ADMIN_CONFIRMATION = "Waiting for admin confirmation"
const TRANSACTION_SHIPPER_SENT = "The package is being delivered by a courier"
const TRANSACTION_SHIPPER_RECEIVED = "Order has been received"

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
	ID            uint           `json:"ID" form:"ID" gorm:"primaryKey"`
	TransactionId uint           `json:"trasaction_id" form:"trasaction_id" gorm:"column:trasaction_id"`
	ShoesId       uint           `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Qty           int            `json:"qty" form:"qty"`
	Price         float64        `json:"price" form:"price"`
	Size          int            `json:"size" form:"size"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type TransactionResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    Transaction `json:"data"`
}

type TransactionListResponse struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Data    []Transaction `json:"data"`
}
