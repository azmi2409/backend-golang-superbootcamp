package routes

import (
	"api-store/controller/admin"
	"api-store/middleware/superadmin"

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

	r.POST("/admin/login", admin.Login)

	adminMiddleware := r.Group("/admin")
	adminMiddleware.Use(superadmin.CheckSuperAdmin())
	{
		adminMiddleware.POST("/register", admin.Register)
	}

	//r.POST("/login", controllers.Login)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API Store",
		})
	})

	return r
}
