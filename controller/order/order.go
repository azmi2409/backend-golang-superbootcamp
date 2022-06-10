package order

import (
	"api-store/middleware"
	"api-store/models"
	"api-store/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, models.NewHttpError(err.Error()))
		return
	}

	//Get user id from token
	tokenString := token.ExtractToken(c)
	id := token.ParseTokenID(tokenString)
	if id == 0 {
		c.JSON(http.StatusUnauthorized, models.NewHttpError("Unauthorized"))
		return
	}

	//Get user cart and cart_items
	var cart models.Cart
	db.Joins("CartItem").Where("user_id = ?", id).First(&cart)
	if cart.ID == 0 {
		c.JSON(http.StatusBadRequest, models.NewHttpError("Cart not found"))
		return
	}

	//return
	c.JSON(http.StatusOK, cart)
}

func OrderRoutes(r *gin.RouterGroup) {
	r.GET("/checkout", middleware.CheckToken, CreateOrder)
}
