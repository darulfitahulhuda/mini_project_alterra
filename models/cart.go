package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	*gorm.Model
	UserId     uint    `json:"user_id" form:"user_id" gorm:"column:user_id"`
	TotalPrice float64 `json:"total_price" form:"total_price" gorm:"column:total_price"`
}

type CartItems struct {
	*gorm.Model
	CartId  uint   `json:"cart_id" form:"cart_id" gorm:"column:cart_id"`
	ShoesId uint   `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Size    int    `json:"size" form:"size"`
	Qty     int    `json:"qty" form:"qty"`
	Status  string `json:"status" form:"status"`
}
