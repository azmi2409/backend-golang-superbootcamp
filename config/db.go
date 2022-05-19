package config

import (
	"beeleaf-store/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := utils.Getenv("DSN", "root:root@tcp(localhost)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
