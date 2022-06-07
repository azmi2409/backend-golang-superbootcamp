package product

import (
	"api-store/middleware/superadmin"
	"api-store/models"
	"api-store/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductInput struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Price       float64  `json:"price" binding:"required"`
	SKU         string   `json:"sku" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	ImageURL    []string `json:"image_url"`
	Slug        string   `json:"slug"`
}

// ShowAccount godoc
// @Summary      Add Product
// @Description  Add Product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object} models.HttpSuccess
// @Failure      400  {object} models.HttpError
// @Failure      404  {object} models.HttpError
// @Failure      500  {object} models.HttpError
// @Router       /product/ [post]
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
		Slug:        utils.CreateSlug(Product.Name),
	}

	db.Create(&product)
	//Insert Picture
	for _, image := range Product.ImageURL {
		productImage := models.ProductImage{
			ProductID: product.ID,
			ImageURL:  image,
		}
		db.Create(&productImage)
	}

	c.JSON(http.StatusCreated, models.NewHttpSuccess("Product added successfully"))

}

// ShowAccount godoc
// @Summary      Update Product
// @Description  Update Product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object} models.HttpSuccess
// @Failure      400  {object} models.HttpError
// @Failure      404  {object} models.HttpError
// @Failure      500  {object} models.HttpError
// @Router       /product/:id [put]
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

// ShowAccount godoc
// @Summary      Get All Product
// @Description  Get All Product
// @Tags         Product
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object} models.HttpSuccess
// @Failure      400  {object} models.HttpError
// @Failure      404  {object} models.HttpError
// @Failure      500  {object} models.HttpError
// @Router       /product [get]
func GetAllProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Joins("Category").Find(&products)
	for i := range products {
		var images []models.ProductImage
		db.Find(&images, "product_id = ?", products[i].ID)
		products[i].ProductImages = images
	}
	var productsJSON []ProductInput
	for _, product := range products {
		var productJSON ProductInput
		productJSON.Name = product.Name
		productJSON.Description = product.Description
		productJSON.Price = product.Price
		productJSON.SKU = product.SKU
		productJSON.Category = product.Category.Name
		productJSON.Slug = product.Slug
		productJSON.ImageURL = []string{}
		for _, image := range product.ProductImages {
			productJSON.ImageURL = append(productJSON.ImageURL, image.ImageURL)
		}
		productsJSON = append(productsJSON, productJSON)
	}

	c.JSON(http.StatusOK, productsJSON)
}

// ShowAccount godoc
// @Summary      Get All Product By Category
// @Description  Get All Product By Category
// @Tags         Product
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object} models.HttpSuccess
// @Failure      400  {object} models.HttpError
// @Failure      404  {object} models.HttpError
// @Failure      500  {object} models.HttpError
// @Router       /product/categories/:id [get]
func GetProductsByCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	category := c.Param("category")

	var categoryID models.Category
	db.Where("name = ?", category).First(&categoryID)

	var products []models.Product
	db.Joins("Category").Where("category_id = ?", categoryID.ID).Find(&products)
	for i := range products {
		var images []models.ProductImage
		db.Find(&images, "product_id = ?", products[i].ID)
		products[i].ProductImages = images
	}

	c.JSON(http.StatusOK, products)
}

// ShowAccount godoc
// @Summary      Get All Product By ID
// @Description  Get All Product By ID
// @Tags         Product
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object} models.HttpSuccess
// @Failure      400  {object} models.HttpError
// @Failure      404  {object} models.HttpError
// @Failure      500  {object} models.HttpError
// @Router       /product/:id [get]
func GetProductByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var product models.Product
	db.Joins("Category").First(&product, id)
	if product.ID == 0 {
		c.JSON(http.StatusBadRequest, models.NewHttpError("Product not found"))
		return
	}
	var images []models.ProductImage
	db.Find(&images, "product_id = ?", product.ID)
	product.ProductImages = images

	c.JSON(http.StatusOK, product)
}

// ShowAccount godoc
// @Summary      Get Delete Product
// @Description  Get Delete Product
// @Tags         Product
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object} models.HttpSuccess
// @Failure      400  {object} models.HttpError
// @Failure      404  {object} models.HttpError
// @Failure      500  {object} models.HttpError
// @Router       /product/:id [delete]
func DeleteProductByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	db.Delete(&models.Product{}, id)
	c.JSON(http.StatusOK, models.NewHttpSuccess("Product deleted successfully"))
}

// ShowAccount godoc
// @Summary      Search Product
// @Description  Search Product
// @Tags         Product
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object} models.HttpSuccess
// @Failure      400  {object} models.HttpError
// @Failure      404  {object} models.HttpError
// @Failure      500  {object} models.HttpError
// @Router       /product/search/ [get]
func SearchProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	query := c.Param("query")
	var products []models.Product
	// put % in query

	db.Joins("Category").Where("name LIKE ?", "%"+query+"%").Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProductBySlug(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	slug := c.Param("slug")
	var product models.Product
	db.Joins("Category").Where("slug = ?", slug).First(&product)
	if product.ID == 0 {
		c.JSON(http.StatusBadRequest, models.NewHttpError("Product not found"))
		return
	}
	var images []models.ProductImage
	db.Find(&images, "product_id = ?", product.ID)
	product.ProductImages = images

	var productJSON ProductInput
	productJSON.Name = product.Name
	productJSON.Description = product.Description
	productJSON.Price = product.Price
	productJSON.SKU = product.SKU
	productJSON.Category = product.Category.Name
	productJSON.Slug = product.Slug
	productJSON.ImageURL = []string{}
	for _, image := range product.ProductImages {
		productJSON.ImageURL = append(productJSON.ImageURL, image.ImageURL)
	}

	c.JSON(http.StatusOK, productJSON)
}

func ProductRoutes(r *gin.RouterGroup) {

	r.GET("/", GetAllProducts)
	r.GET("/:slug", GetProductBySlug)
	r.GET("/category/:category", GetProductsByCategory)
	r.GET("/search/:query", SearchProduct)

	//add auth middleware
	r.Use(superadmin.CheckSuperAdmin())
	r.POST("/", AddProduct)
	r.PUT("/:id", UpdateProduct)
	r.DELETE("/:id", DeleteProductByID)
}
