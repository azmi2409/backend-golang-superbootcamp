package admin

import (
	"api-store/middleware/superadmin"
	"api-store/models"
	"api-store/utils"
	"api-store/utils/storage"
	"api-store/utils/token"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminInput struct {
	Name       string `json:"name"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	SuperAdmin bool   `json:"superAdmin" default:"false"`
}

type ImageInput struct {
	Image []byte `json:"image" binding:"required"`
	Name  string `json:"name"`
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
	newAdmin.SuperAdmin = Admin.SuperAdmin
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
	var Admintype string

	if checkAdmin.SuperAdmin {
		Admintype = "superAdmin"
	} else {
		Admintype = "admin"
	}

	var data = map[string]interface{}{
		"id":   checkAdmin.ID,
		"name": checkAdmin.Name,
		"type": Admintype,
	}

	token, err := token.GenerateToken(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output = map[string]interface{}{
		"id":    checkAdmin.ID,
		"name":  checkAdmin.Name,
		"type":  Admintype,
		"token": token,
	}

	c.JSON(http.StatusOK, output)
}

func UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := header.Filename

	img, _ := ioutil.ReadAll(file)

	// fmt.Println(string(img))
	if !utils.CheckImageExtension(fileName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image extension is not valid"})
		return
	}

	//Get Image Extension

	//Save Image to Storage
	path, err := storage.UploadFiles(img, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"path": path})

}

func UploadImageBase64(c *gin.Context) {
	var Image ImageInput
	if err := c.ShouldBindJSON(&Image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Save Image to Storage

	path, err := storage.UploadBase64(Image.Image, Image.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"path": path})

}

func AdminRoutes(r *gin.RouterGroup) {
	r.POST("/login", Login)

	//Use Auth
	r.Use(superadmin.CheckSuperAdmin())
	r.POST("/register", Register)
	r.POST("/upload", UploadImage)
	r.POST("/upload-base", UploadImageBase64)
}
