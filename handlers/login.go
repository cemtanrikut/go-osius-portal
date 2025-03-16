package handlers

import (
	"net/http"

	"main.go/config"
	"main.go/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 📌 **JWT için Secret Key**
var jwtSecret = []byte("supersecretkey")

// 📌 **Login Request Modeli**
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// // 📌 **JWT Token Oluşturma Fonksiyonu**
// func generateJWT(userEmail string, userType string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"email": userEmail,
// 		"type":  userType,                              // "worker" veya "customer" olarak belirtiyoruz
// 		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 24 saat geçerli olacak
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(jwtSecret)
// }

// 📌 **Login Fonksiyonu**
func Login(c *gin.Context) {
	var req LoginRequest

	// 📌 **Request'ten Email ve Password Al**
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Eksik veya hatalı giriş!"})
		return
	}

	var worker models.Worker
	var customer models.Customer

	// 📌 **Önce Workers tablosunda arama yap**
	if err := config.DB.Where("email = ?", req.Email).First(&worker).Error; err == nil {
		// 🎯 **Worker bulundu, şifreyi kontrol et**
		if err := bcrypt.CompareHashAndPassword([]byte(worker.Password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre yanlış!"})
			return
		}

		// 📌 **JWT Token oluştur ve döndür**
		// token, _ := generateJWT(worker.Email, "worker")
		// c.JSON(http.StatusOK, gin.H{"token": token, "userType": "worker"})
		return
	}

	// 📌 **Worker bulunamadıysa, Customers tablosuna bak**
	if err := config.DB.Where("email = ?", req.Email).First(&customer).Error; err == nil {
		// 🎯 **Customer bulundu, şifreyi kontrol et**
		if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre yanlış!"})
			return
		}

		// 📌 **JWT Token oluştur ve döndür**
		// token, _ := generateJWT(customer.Email, "customer")
		// c.JSON(http.StatusOK, gin.H{"token": token, "userType": "customer"})
		return
	}

	// 📌 **Eğer hiçbir kullanıcı bulunamazsa hata dön**
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı!"})
}
