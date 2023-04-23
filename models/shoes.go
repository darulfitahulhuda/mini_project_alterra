package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Shoes struct {
	*gorm.Model
	Name   string         `json:"name" form:"name"`
	Images pq.StringArray `json:"images" form:"images" gorm:"type:text[]"`
	Price  float64        `json:"price" form:"price"`
	Gender string         `json:"gender" form:"gender"`
}
