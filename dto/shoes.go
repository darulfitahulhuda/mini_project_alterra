package dto

type Shoes struct {
	Name        string      `json:"name"`
	Images      []string    `json:"images"`
	Price       float64     `json:"price"`
	Gender      string      `json:"gender"`
	Description string      `json:"description"`
	Sizes       []ShoesSize `json:"sizes"`
	Brand       string      `json:"brand"`
}
type ShoesSize struct {
	Size    int `json:"size"`
	Qty     int `json:"qty"`
	ShoesId int `json:"shoes_id"`
}
