package routes

import (
	"api-store/controller/admin"
	"api-store/controller/product"
	"api-store/controller/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Api Main Route
	api := r.Group("/api/v1")

	//Admin Routes
	adminRoutes := api.Group("/admin")
	userRoutes := api.Group("/user")
	productRoutes := api.Group("/product")
	//	orderRoutes := api.Group("/order")
	//	checkoutRoutes := api.Group("/checkout")

	admin.AdminRoutes(adminRoutes)
	user.UserRoutes(userRoutes)
	product.ProductRoutes(productRoutes)

	return r
}
