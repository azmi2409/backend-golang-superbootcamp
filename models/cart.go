package models

import "time"

type Cart struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
