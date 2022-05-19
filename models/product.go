package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	SKU         string    `json:"sku"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
