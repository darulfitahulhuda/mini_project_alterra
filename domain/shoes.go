package domain

import "gorm.io/gorm"

type Shoes struct {
	*gorm.Model
	Name   string   `json:"name" form:"name"`
	Images []string `json:"images" form:"images"`
	Price  float64  `json:"price" form:"price"`
	Gender string   `json:"gender" form:"gender"`
}
