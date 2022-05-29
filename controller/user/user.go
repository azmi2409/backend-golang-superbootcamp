package user

import (
	"api-store/models"
	"api-store/utils"
	"api-store/utils/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ShowAccount godoc
// @Summary      Register User
// @Description  Register User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  "User registered successfully"
// @Router       /user/register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var User models.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check If Email Valid
	email := User.Email
	if !utils.ValidateEmail(*email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is not valid"})
		return
	}

	//Check if User already Registered
	var checkUser models.User
	db.Where("email = ?", User.Email).First(&checkUser)
	if checkUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	//Hash Password
	hashedPassword, err := utils.Encrypt(User.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Create User
	var newUser models.User
	newUser.Name = User.Name
	newUser.Email = User.Email
	newUser.Age = User.Age
	//	newUser.Birthday = User.Birthday
	newUser.City = User.City
	newUser.Country = User.Country
	newUser.Address = User.Address
	newUser.Phone = User.Phone
	newUser.ZipCode = User.ZipCode
	newUser.Password = hashedPassword
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	db.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var UserInput UserInput
	if err := c.ShouldBindJSON(&UserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check if User already Registered
	var checkUser models.User
	db.Where("email = ?", UserInput.Email).First(&checkUser)
	if checkUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not registered"})
		return
	}

	//Check if Password is Correct
	if !utils.CheckPasswordHash(UserInput.Password, checkUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is not correct"})
		return
	}

	//Generate JSON
	var UserData = map[string]interface{}{
		"id":    checkUser.ID,
		"name":  checkUser.Name,
		"email": checkUser.Email,
	}

	token, err := token.GenerateToken(UserData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Successfully", "token": token})
}

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/register", Register)
	r.POST("/login", Login)

}
