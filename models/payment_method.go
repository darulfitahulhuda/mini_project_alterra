package models

import "gorm.io/gorm"

type PaymentMethod struct {
	*gorm.Model
	Name   string `json:"name" form:"name"`
	Status string `json:"status" form:"status"`
}
