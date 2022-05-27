package routes

import (
	"api-store/controller/admin"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// swagger embed files
	// gin-swagger middleware
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

	//Api Main Route
	api := r.Group("/api/v1")

	//Admin Routes
	adminRoutes := api.Group("/admin")
	//	userRoutes := api.Group("/user")
	//	productRoutes := api.Group("/product")
	//	orderRoutes := api.Group("/order")
	//	checkoutRoutes := api.Group("/checkout")

	admin.AdminRoutes(adminRoutes)

	return r
}
