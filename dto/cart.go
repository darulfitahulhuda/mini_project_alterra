package dto

type Cart struct {
	ShoesId int    `json:"shoes_id" form:"shoes_id"`
	Size    int    `json:"size" form:"size"`
	Qty     int    `json:"qty" form:"qty"`
	Status  string `json:"status" form:"status"`
}
