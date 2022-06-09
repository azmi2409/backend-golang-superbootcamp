package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	SKU         string  `json:"sku" gorm:"unique"`
	CategoryID  uint    `json:"category_id"`
	Slug        string  `json:"slug"`

	Category      Category       `gorm:"foreignkey:CategoryID"`
	ProductImages []ProductImage `gorm:"foreignkey:ProductID"`
}
