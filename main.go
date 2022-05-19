package main

import (
	"api-store/config"
	"fmt"
)

func main() {
	//programmatically set swagger info

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
