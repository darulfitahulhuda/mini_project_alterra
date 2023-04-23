package dto

import "time"

type CreateUser struct {
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Address     string     `json:"address"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
