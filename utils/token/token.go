package token

import (
	"api-store/utils"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var API_SECRET = utils.GetEnv("API_SECRET", "supersecret")

func GenerateToken(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token.SignedString([]byte(API_SECRET))
}

func TokenValid(tokenString string) (uint, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return 0, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["id"]), 10, 32)
		if err != nil {
			return 0, false
		}
		return uint(uid), true
	}
	return 0, false
}

func ExtractToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		return ""
	}
	return strings.Replace(token, "Bearer ", "", -1)
}

func ExtractTokenID(c *gin.Context) uint {
	token := ExtractToken(c)
	if len(token) == 0 {
		return 0
	}
	uid, _ := TokenValid(token)
	return uid
}
