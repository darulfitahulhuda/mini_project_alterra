package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipping struct {
	*gorm.Model
	TransactionId uint       `json:"trasaction_id" form:"trasaction_id" gorm:"column:trasaction_id"`
	Address       string     `json:"address" form:"address"`
	Price         float64    `json:"price" form:"price"`
	Method        string     `json:"method" form:"method"`
	DeliveriDate  *time.Time `json:"delivery_date" form:"delivery_date" gorm:"column:delivery_date"`
}
