package main

import (
	"api-store/config"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	//programmatically set swagger info

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	db := config.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	//check db connection
	if db.Error != nil {
		panic(db.Error)
	}

	//check db tables
	fmt.Printf("Db Connection Successful\n")
}
