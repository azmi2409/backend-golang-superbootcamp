package main

import (
	"api-store/config"
	"api-store/routes"
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

	//setup storage
	//storage.SetupStorage()

	//check db tables
	fmt.Printf("Db Connection Successful\n")

	r := routes.SetupRouter(db)
	r.Run(":8080")
}
