package product

import (
	"api-store/middleware/superadmin"
	"api-store/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	SKU         string  `json:"sku" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	ImageURL    string  `json:"image_url"`
}

func AddProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Product ProductInput
	if err := c.ShouldBindJSON(&Product); err != nil {
		c.JSON(http.StatusBadRequest, models.NewHttpError(err.Error()))
		return
	}

	//Check if category exist
	var category models.Category
	db.Where("name = ?", Product.Category).First(&category)
	if category.ID == 0 {
		//create category
		category = models.Category{
			Name:        Product.Category,
			Description: "",
		}
		db.Create(&category)
	}

	db.Where("name = ?", Product.Category).First(&category)

	//Convert product to model
	product := models.Product{
		Name:        Product.Name,
		Description: Product.Description,
		Price:       Product.Price,
		SKU:         Product.SKU,
		CategoryID:  category.ID,
	}

	db.Create(&product)
	//Insert Picture
	if Product.ImageURL != "" {
		db.Create(&models.ProductImage{
			ProductID: product.ID,
			ImageURL:  Product.ImageURL,
		})
	}

	c.JSON(http.StatusOK, models.NewHttpSuccess("Product added successfully"))

}

func UpdateProduct(c *gin.Context) {
	//get ID params
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	//Check if product exist
	var product models.Product
	db.Where("id = ?", id).First(&product)
	if product.ID == 0 {
		c.JSON(http.StatusBadRequest, models.NewHttpError("Product not found"))
		return
	}

	var ProductInput ProductInput
	if err := c.ShouldBindJSON(&ProductInput); err != nil {
		c.JSON(http.StatusBadRequest, models.NewHttpError(err.Error()))
		return
	}
	//update product
	product.Name = ProductInput.Name
	product.Description = ProductInput.Description
	product.Price = ProductInput.Price
	product.SKU = ProductInput.SKU
	db.Save(&product)
	c.JSON(http.StatusOK, models.NewHttpSuccess("Product updated successfully"))
}

func GetAllProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Joins("Category").Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProductsByCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	category := c.Param("category")

	//get Category ID
	var categoryModel models.Category
	db.Where("name = ?", category).First(&categoryModel)
	if categoryModel.ID == 0 {
		c.JSON(http.StatusBadRequest, models.NewHttpError("Category not found"))
		return
	}
	var products []models.Product
	db.Joins("Category").Where("category_id = ?", categoryModel.ID).Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var product models.Product
	db.Joins("Category").First(&product, id)
	if product.ID == 0 {
		c.JSON(http.StatusBadRequest, models.NewHttpError("Product not found"))
		return
	}
	c.JSON(http.StatusOK, product)
}

func DeleteProductByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	db.Delete(&models.Product{}, id)
	c.JSON(http.StatusOK, models.NewHttpSuccess("Product deleted successfully"))
}

func SearchProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	query := c.Param("query")
	var products []models.Product
	// put % in query

	db.Joins("Category").Where("name LIKE ?", "%"+query+"%").Find(&products)
	c.JSON(http.StatusOK, products)
}

func ProductRoutes(r *gin.RouterGroup) {

	r.GET("/", GetAllProducts)
	r.GET("/:id", GetProductByID)
	r.GET("/category/:category", GetProductsByCategory)
	r.GET("/search/:query", SearchProduct)

	//add auth middleware
	r.Use(superadmin.CheckSuperAdmin())
	r.POST("/", AddProduct)
	r.PUT("/:id", UpdateProduct)
	r.DELETE("/:id", DeleteProductByID)
}
