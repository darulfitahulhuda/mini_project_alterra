package models

import (
	"gorm.io/gorm"
)

type CartItems struct {
	*gorm.Model
	CartId  uint   `json:"cart_id" form:"cart_id" gorm:"column:cart_id"`
	ShoesId uint   `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Qty     int    `json:"qty" form:"qty"`
	Status  string `json:"status" form:"status"`
}
