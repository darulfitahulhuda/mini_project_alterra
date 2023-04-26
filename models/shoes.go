package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Shoes struct {
	*gorm.Model
	Name        string         `json:"name" form:"name"`
	Images      pq.StringArray `json:"images" form:"images" gorm:"type:text[]"`
	Price       float64        `json:"price" form:"price"`
	Gender      string         `json:"gender" form:"gender"`
	ShoesDetail ShoesDetail    `json:"detail" form:"detail" gorm:"foreignkey:shoes_id"`
}

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

type ShoesListResponse struct {
	Message string          `json:"message"`
	Status  int             `json:"status"`
	Data    []ShoesListData `json:"data"`
}
type ShoesListData struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Images []string `json:"images"`
	Price  float64  `json:"price"`
	Gender string   `json:"gender"`
}
type ShoesDetailResponse struct {
	Message string          `json:"message"`
	Status  int             `json:"status"`
	Data    ShoesDetailData `json:"data"`
}

type ShoesDetailData struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Images      []string `json:"images"`
	Price       float64  `json:"price"`
	Gender      string   `json:"gender"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Color       []string `json:"color" `
	Size        []int64  `json:"size" `
	Qty         int      `json:"qty" `
	Brand       string   `json:"brand"`
}
