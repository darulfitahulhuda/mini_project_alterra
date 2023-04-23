package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ShoesDetail struct {
	*gorm.Model
	ShoesId     uint           `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Description string         `json:"description" form:"description"`
	Category    string         `json:"category" form:"category"`
	Color       pq.StringArray `json:"color" form:"color" gorm:"type:text[]"`
	Size        pq.Int64Array  `json:"size" form:"size" gorm:"type:integer[]"`
	Qty         int            `json:"qty" form:"qty"`
	Brand       string         `json:"brand" form:"brand"`
}
