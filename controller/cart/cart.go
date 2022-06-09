package cart

import (
	"api-store/middleware"
	"api-store/models"
	"api-store/utils/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartInput struct {
	SKU      string `json:"sku" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type CartItemOutput struct {
	SKU         string `json:"sku"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

func AddtoCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cart CartInput
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, models.NewHttpError(err.Error()))
		return
	}

	//Check if product exist
	var product models.Product
	db.Where("sku = ?", cart.SKU).First(&product)
	if product.ID == 0 {
		c.JSON(http.StatusBadRequest, models.NewHttpError("Product not found"))
		return
	}

	tokenString := token.ExtractToken(c)
	id := token.ParseTokenID(tokenString)
	if id == 0 {
		c.JSON(http.StatusUnauthorized, models.NewHttpError("Unauthorized"))
		return
	}

	//Check if cart exist
	var cartDB models.Cart
	db.Where("user_id = ?", id).First(&cartDB)
	if cartDB.ID == 0 {
		cartDB.UserID = id
		db.Create(&cartDB)
	}

	//Check if product exist in cart
	var cartItem models.CartItem
	db.Where("cart_id = ? AND product_id = ?", cartDB.ID, product.ID).First(&cartItem)
	if cartItem.ID == 0 {
		//input to cart_items
		var cartItem models.CartItem
		cartItem.CartID = cartDB.ID
		cartItem.ProductID = product.ID
		cartItem.Quantity = cart.Quantity
		cartItem.CreatedAt = time.Now()
		db.Create(&cartItem)
	} else {
		cartItem.Quantity = cartItem.Quantity + cart.Quantity
		db.Save(&cartItem)
	}
	cartDB.Total += 1
	db.Save(&cartDB)

	c.JSON(http.StatusOK, models.NewHttpSuccess("Product added to cart successfully"))
}

func ViewCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	tokenString := token.ExtractToken(c)
	id := token.ParseTokenID(tokenString)
	if id == 0 {
		c.JSON(http.StatusUnauthorized, models.NewHttpError("Unauthorized"))
		return
	}

	//Check if cart exist
	var cartDB models.Cart
	db.Where("user_id = ?", id).First(&cartDB)
	if cartDB.ID == 0 {
		c.JSON(http.StatusOK, []CartItemOutput{})
		return
	}

	//Get cart items
	var cartItems []models.CartItem
	db.Joins("Product").Where("cart_id = ?", cartDB.ID).Find(&cartItems)
	var cartItemsOutput []CartItemOutput
	for _, cartItem := range cartItems {
		cartItemOutput := CartItemOutput{
			SKU:         cartItem.Product.SKU,
			ProductName: cartItem.Product.Name,
			Quantity:    cartItem.Quantity,
		}
		cartItemsOutput = append(cartItemsOutput, cartItemOutput)
	}

	c.JSON(http.StatusOK, cartItemsOutput)
}

func CartRoutes(r *gin.RouterGroup) {
	r.POST("/", middleware.CheckToken, AddtoCart)
	r.GET("/", middleware.CheckToken, ViewCart)
}
