package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryJSON struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func GetAllCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories []CategoryJSON
	db.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func CategoryRoute(r *gin.RouterGroup) {
	r.GET("/", GetAllCategories)
}
