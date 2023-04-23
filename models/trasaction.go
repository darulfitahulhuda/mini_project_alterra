package models

import "gorm.io/gorm"

type Transaction struct {
	*gorm.Model
	UserId          uint    `json:"user_id" form:"user_id" gorm:"column:user_id"`
	PaymentMethodId uint    `json:"payment_method_id" form:"payment_method_id" gorm:"column:payment_method_id"`
	TotalPrice      float64 `json:"total_price" form:"total_price" gorm:"column:total_price"`
	Status          string  `json:"status" form:"status"`
}
