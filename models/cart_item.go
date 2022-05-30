package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`

	Product Product `gorm:"foreignkey:ProductID"`
}
