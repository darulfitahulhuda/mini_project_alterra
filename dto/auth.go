package dto

type CreateToken struct {
	Id       int    `json:"id"`
	UserType string `json:"user_type" form:"user_type" `
}
