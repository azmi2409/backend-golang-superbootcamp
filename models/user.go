package models

import (
	"time"
)

type User struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Email     *string    `json:"email"`
	Password  string     `json:"password"`
	Address   string     `json:"address"`
	Age       uint       `json:"age"`
	Birthday  *time.Time `json:"birthday"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	Cart []Cart `json:"-"`
}
