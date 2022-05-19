package config

import (
	"api-store/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	user := "postgres"
	password := "Azmi2409Revo2014"
	host := "db.cwrfdvnvvcedqjylgvms.supabase.co"
	port := "5432"
	dbname := "postgres"
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{})

	return db
}
