package models

import (
	"main/dto"

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
	Sizes       []ShoesSize    `json:"sizes" form:"sizes" gorm:"foreignkey:shoes_id"`
}

type ShoesDetail struct {
	*gorm.Model
	ShoesId     uint   `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Description string `json:"description" form:"description"`
	Brand       string `json:"brand" form:"brand"`
}

type ShoesSize struct {
	*gorm.Model
	ShoesId uint `json:"shoes_id" form:"shoes_id" gorm:"column:shoes_id"`
	Size    int  `json:"size" form:"size"`
	Qty     int  `json:"qty" form:"qty"`
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
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Images      []string        `json:"images"`
	Price       float64         `json:"price"`
	Gender      string          `json:"gender"`
	Description string          `json:"description"`
	Brand       string          `json:"brand"`
	Sizes       []dto.ShoesSize `json:"sizes"`
}
