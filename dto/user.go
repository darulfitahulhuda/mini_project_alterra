package dto

import "time"

type CreateUser struct {
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Address     string     `json:"address"`
	UserType    string     `json:"user_type" form:"user_type" default:"user"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type" form:"user_type" default:"user"`
}

type UpdateUser struct {
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Address     string     `json:"address"`
}
