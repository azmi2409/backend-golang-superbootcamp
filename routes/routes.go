package routes

import (
	"api-store/controller/admin"
	"api-store/controller/cart"
	"api-store/controller/categories"
	"api-store/controller/product"
	"api-store/controller/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

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

	userRoutes := api.Group("/user")
	productRoutes := api.Group("/product")
	cartRoutes := api.Group("/cart")
	categoryRoutes := api.Group("/category")
	//	orderRoutes := api.Group("/order")
	//	checkoutRoutes := api.Group("/checkout")

	admin.AdminRoutes(adminRoutes)
	product.ProductRoutes(admProduct)

	user.UserRoutes(userRoutes)
	product.ProductRoutes(productRoutes)
	cart.CartRoutes(cartRoutes)
	categories.CategoryRoute(categoryRoutes)

	return r
}
