package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint `json:"user_id"`
	Total  int  `json:"total" default:"0"`

	User     User       `gorm:"foreignkey:UserID"`
	CartItem []CartItem `gorm:"foreignkey:CartID"`
}
