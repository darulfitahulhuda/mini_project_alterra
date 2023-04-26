package dto

type Shoes struct {
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
