package models

import "gorm.io/gorm"

type PaymentMethod struct {
	*gorm.Model
	TransactionId uint           `json:"transaction_id" form:"transaction_id" gorm:"column:transaction_id"`
	Name          string         `json:"name" form:"name"`
	Status        string         `json:"status" form:"status"`
	CodePayment   string         `json:"code_payment" form:"code_payment" gorm:"column:code_payment"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type PaymentMethodResponse struct {
	Name        string `json:"name" form:"name"`
	Status      string `json:"status" form:"status"`
	CodePayment string `json:"code_payment" form:"code_payment" gorm:"column:code_payment"`
}
