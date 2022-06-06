package categories

import (
	"api-store/models"
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
	var categories []models.Category
	db.Find(&categories)

	var categoriesJSON []CategoryJSON
	for _, category := range categories {
		categoriesJSON = append(categoriesJSON, CategoryJSON{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ImageURL:    category.Image_url,
		})
	}

	c.JSON(http.StatusOK, categoriesJSON)
}

func CategoryRoute(r *gin.RouterGroup) {
	r.GET("/", GetAllCategories)
}
