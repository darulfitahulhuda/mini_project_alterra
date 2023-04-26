package dto

import "time"

type Transaction struct {
	UserId        int                 `json:"user_id" form:"user_id"`
	TotalPrice    float64             `json:"total_price" form:"total_price"`
	Status        string              `json:"status" form:"status"`
	NamePayment   string              `json:"name_payment" form:"name_payment"`
	StatusPayment string              `json:"status_payment" form:"status_payment"`
	CodePayment   string              `json:"code_payment" form:"code_payment" gorm:"column:code_payment"`
	Products      []TransactionItems  `json:"products" form:"products"`
	Shipping      TransactionShipping `json:"shipping" form:"shipping"`
}

type TransactionItems struct {
	ShoesId int     `json:"shoes_id" form:"shoes_id"`
	Qty     int     `json:"qty" form:"qty"`
	Price   float64 `json:"price" form:"price"`
}

type TransactionShipping struct {
	Address      string     `json:"address" form:"address"`
	Price        float32    `json:"price" form:"price"`
	Method       string     `json:"method" form:"method"`
	DeliveriDate *time.Time `json:"delivery_date" form:"delivery_date"`
}
