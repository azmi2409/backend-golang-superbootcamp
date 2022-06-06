package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"image_url"`

	Product []Product `gorm:"foreignkey:CategoryID"`
}
