package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string     `json:"name"`
	Email    *string    `json:"email"`
	Password string     `json:"password"`
	Address  string     `json:"address"`
	Age      uint       `json:"age"`
	Birthday *time.Time `json:"birthday"`
}
