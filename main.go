package main

import (
	"api-store/config"
	"api-store/docs"
	"api-store/routes"
	"api-store/utils"
	"fmt"

	"github.com/joho/godotenv"
)

// @title           FinalProject GO API
// @version         0.1
// @description     This is a simple E-Commerce API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Azmi
// @contact.url    https://www.azmi.web.id
// @contact.email  me@azmi.web.id

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BearerToken

func main() {
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "FinalProject GO API"
	docs.SwaggerInfo.Description = "Simple E-Commerce API"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = "http://backend-final-beeleaf.herokuapp.com"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	enviroment := utils.GetEnv("ENVIROMENT", "development")

	if enviroment == "development" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(err)
		}
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

	port := utils.GetEnv("PORT", "8080")

	r := routes.SetupRouter(db)
	r.Run(":" + port)
}
