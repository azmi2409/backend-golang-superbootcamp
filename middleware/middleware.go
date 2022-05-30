package middleware

import (
	"api-store/models"
	"api-store/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckToken(c *gin.Context) {
	tokenString := token.ExtractToken(c)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		c.Abort()
		return
	}
	id := token.ParseTokenID(tokenString)
	if id != 0 {
		//verify id
		db := c.MustGet("db").(*gorm.DB)
		var user models.User
		db.First(&user, id)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()

}
