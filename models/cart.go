package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	*gorm.Model
	UserId     uint    `json:"user_id" form:"user_id" gorm:"column:user_id"`
	TotalPrice float64 `json:"total_price" form:"total_price" gorm:"column:total_price"`
}
