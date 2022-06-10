package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	Total     float64 `json:"total"`
	PaymentID uint    `json:"payment_id"`
	Status    string  `json:"status"`

	User      User        `gorm:"foreignkey:UserID"`
	OrderItem []OrderItem `gorm:"foreignkey:OrderID"`
}
