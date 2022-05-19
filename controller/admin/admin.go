package admin

import (
	"net/http"
	"time"

	"api-store/models"

	"api-store/utils"
	"api-store/utils/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Admin AdminInput
	if err := c.ShouldBindJSON(&Admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check If Email Valid
	if !utils.ValidateEmail(Admin.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is not valid"})
		return
	}

	//Check if Admin already Registered
	var checkAdmin models.Admin
	db.Where("email = ?", Admin.Email).First(&checkAdmin)
	if checkAdmin.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Admin already registered"})
		return
	}

	//Hash Password
	hashedPassword, err := utils.Encrypt(Admin.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Create Admin
	var newAdmin models.Admin
	newAdmin.Name = Admin.Name
	newAdmin.Email = &Admin.Email
	newAdmin.Password = hashedPassword
	newAdmin.CreatedAt = time.Now()
	newAdmin.UpdatedAt = time.Now()
	db.Create(&newAdmin)

	c.JSON(http.StatusOK, gin.H{"message": "Admin registered successfully"})
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Admin AdminInput
	if err := c.ShouldBindJSON(&Admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check If Email Valid
	if !utils.ValidateEmail(Admin.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is not valid"})
		return
	}

	//Check if Admin already Registered
	var checkAdmin models.Admin
	db.Where("email = ?", Admin.Email).First(&checkAdmin)
	if checkAdmin.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Admin not registered"})
		return
	}

	//Check if Password Valid
	if !utils.CheckPasswordHash(Admin.Password, checkAdmin.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is not valid"})
		return
	}

	//Generate JWT Token
	token, err := token.GenerateToken(checkAdmin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
