package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`

	Order   Order   `gorm:"foreignkey:OrderID"`
	Product Product `gorm:"foreignkey:ProductID"`
}
