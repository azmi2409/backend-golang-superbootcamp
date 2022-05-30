package config

import (
	"api-store/models"
	"api-store/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	user := utils.GetEnv("DB_USERNAME", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "root")
	host := utils.GetEnv("DB_HOST", "localhost")
	port := utils.GetEnv("DB_PORT", "5432")
	dbname := utils.GetEnv("DB_NAME", "postgres")
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Admin{}, &models.User{}, &models.Product{}, &models.Category{}, &models.Order{}, &models.OrderItem{}, &models.ProductImage{}, &models.Cart{}, &models.CartItem{})

	return db
}
