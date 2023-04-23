package models

import "gorm.io/gorm"

type TransactionDetail struct {
	*gorm.Model
	TransactionId uint    `json:"trasaction_id" form:"trasaction_id" gorm:"column:trasaction_id"`
	ShoesId       uint    `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Qty           int     `json:"qty" form:"qty"`
	Price         float64 `json:"price" form:"price"`
}
