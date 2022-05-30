package routes

import (
	"api-store/controller/admin"
	"api-store/controller/cart"
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

	config := cors.New(corsConfig)

	r.Use(config)

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Api Main Route
	api := r.Group("/api/v1")

	//Admin Routes
	adminRoutes := api.Group("/admin")
	admProduct := adminRoutes.Group("/product")
	admProduct.Use(config)

	userRoutes := api.Group("/user")
	productRoutes := api.Group("/product")
	cartRoutes := api.Group("/cart")
	//	orderRoutes := api.Group("/order")
	//	checkoutRoutes := api.Group("/checkout")

	admin.AdminRoutes(adminRoutes)
	product.ProductRoutes(admProduct)

	user.UserRoutes(userRoutes)
	product.ProductRoutes(productRoutes)
	cart.CartRoutes(cartRoutes)

	return r
}
