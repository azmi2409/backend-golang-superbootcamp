package utils

import (
	"net/mail"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Decrypt(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidateEmail(email string) bool {
	//validate email regex
	_, err := mail.ParseAddress(email)
	return err == nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckImageExtension(fileName string) bool {
	arrayOfExt := []string{"jpg", "jpeg", "png", "gif", "bmp", "webp", "svg", "ico", "tiff", "tif"}
	ext := filepath.Ext(fileName)
	for _, v := range arrayOfExt {
		if ext == "."+v {
			return true
		}
	}
	return false
}
