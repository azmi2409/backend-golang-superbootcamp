package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
	Age      uint    `json:"age"`
	//	Birthday *time.Time `json:"birthday"`
	City    string `json:"city"`
	Country string `json:"country"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	ZipCode string `json:"zipcode"`
}
