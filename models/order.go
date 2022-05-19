package models

import "time"

type Order_Detail struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	Total     float64   `json:"total"`
	PaymentID uint      `json:"payment_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order_Item struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	OrderID   uint      `json:"order_id"`
	ProductID uint      `json:"product_id"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
