package models

import (
	"gorm.io/gorm"
)

type Carts struct {
	*gorm.Model
	UserId  uint   `json:"user_id" form:"user_id" gorm:"column:user_id"`
	ShoesId uint   `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Size    int    `json:"size" form:"size"`
	Qty     int    `json:"qty" form:"qty"`
	Status  string `json:"status" form:"status"`
	Shoes   Shoes  `json:"shoes" form:"shoes" gorm:"foreignKey:ID;references:shoes_id"`
}

type CartResponse struct {
	ID         int      `json:"id"`
	ShoesId    int      `json:"shoes_id"`
	Size       int      `json:"size" form:"size"`
	Qty        int      `json:"qty" form:"qty"`
	Price      float64  `json:"price"`
	TotalPrice float64  `json:"total_price"`
	Status     string   `json:"status" form:"status"`
	Name       string   `json:"name"`
	Images     []string `json:"images"`
}
