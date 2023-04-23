package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name        string    `json:"name" form:"name"`
	Email       string    `json:"email" form:"email"`
	Password    string    `json:"password" form:"password"`
	DateOfBirth time.Time `json:"dateOfBirth" form:"dateOfBirth"`
	UserType    string    `json:"userType" form:"userType"`
	Address     string    `json:"address" form:"address"`
}
