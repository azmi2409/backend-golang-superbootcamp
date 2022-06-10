package order

import (
	"api-store/middleware"
	"api-store/models"
	"api-store/utils"
	"api-store/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get user id from token
	tokenString := token.ExtractToken(c)
	id := token.ParseTokenID(tokenString)
	if id == 0 {
		c.JSON(http.StatusUnauthorized, models.NewHttpError("Unauthorized"))
		return
	}

	//Get user cart and cart_items
	var cart models.Cart
	db.Where("user_id = ?", id).First(&cart)
	if cart.ID == 0 {
		c.JSON(http.StatusBadRequest, models.NewHttpError("Cart not found"))
		return
	}
	var items []models.CartItem
	db.Joins("Product").Where("cart_id = ?", cart.ID).Find(&items)

	cart.CartItem = items

	//Calculate total
	total := 0
	for _, v := range cart.CartItem {
		price := int(v.Product.Price)
		total += price * v.Quantity
	}

	//Create Order
	var order models.Order
	order.UserID = id
	order.Total = total
	intId := int(id)
	order.PaymentID = utils.CreateInvoiceNumber(intId)
	order.Status = "Pending"
	db.Create(&order)

	//Create Order Items
	for _, v := range cart.CartItem {
		var orderItem models.OrderItem
		orderItem.OrderID = order.ID
		orderItem.ProductID = v.Product.ID
		orderItem.Quantity = v.Quantity
		db.Create(&orderItem)
	}

	//Remove Cart Items
	db.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})

	c.JSON(http.StatusOK, order)
}

func OrderRoutes(r *gin.RouterGroup) {
	r.GET("/checkout", middleware.CheckToken, CreateOrder)
}
