package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name        string     `json:"name" form:"name"`
	Email       string     `json:"email" form:"email"`
	Password    string     `json:"password" form:"password"`
	DateOfBirth *time.Time `json:"date_of_birth" form:"date_of_birth" gorm:"column:date_of_birth"`
	UserType    string     `json:"user_type" form:"user_type" default:"user" gorm:"column:user_type; default:user"`
	Address     string     `json:"address" form:"address"`
}

const Admin_Type = "admin"
const User_Type = "user"
