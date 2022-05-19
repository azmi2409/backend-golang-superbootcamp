package superadmin

import (
	"api-store/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := token.ExtractToken(c)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}
		err := token.SuperAdminTokenValid(tokenString)
		if err {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	}
}
