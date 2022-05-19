package models

import "time"

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	OrderID   uint      `json:"order_id"`
	ProductID uint      `json:"product_id"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Product Product `json:"-"`
	Order   Order   `json:"-"`
}
