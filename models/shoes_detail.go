package models

import "gorm.io/gorm"

type ShoesDetail struct {
	*gorm.Model
	ShoesId     uint     `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Description string   `json:"description" form:"description"`
	Category    string   `json:"category" form:"category"`
	Color       []string `json:"color" form:"color"`
	Size        []int    `json:"size" form:"size"`
	Qty         int      `json:"qty" form:"qty"`
	Brand       string   `json:"brand" form:"brand"`
}
