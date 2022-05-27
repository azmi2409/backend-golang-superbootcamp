package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`

	User    User      `gorm:"foreignkey:UserID"`
	Product []Product `gorm:"foreignkey:ProductID"`
}
