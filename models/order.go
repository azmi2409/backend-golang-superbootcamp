package models

import "time"

type Order struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	Total     float64   `json:"total"`
	PaymentID uint      `json:"payment_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User  User        `json:"-"`
	Items []OrderItem `json:"-"`
}