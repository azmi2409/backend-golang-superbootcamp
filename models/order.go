package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	Total     int    `json:"total"`
	PaymentID string `json:"payment_id" gorm:"not null,unique_index"`
	Status    string `json:"status"`

	User      User        `gorm:"foreignkey:UserID"`
	OrderItem []OrderItem `gorm:"foreignkey:OrderID"`
}
