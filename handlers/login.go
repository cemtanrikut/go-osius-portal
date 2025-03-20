package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
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

// // 📌 **Login Fonksiyonu**
// func Login(c *gin.Context) {
// 	var req LoginRequest

// 	// 📌 **Request'ten Email ve Password Al**
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Eksik veya hatalı giriş!"})
// 		return
// 	}

// 	var worker models.Worker
// 	var customer models.Customer

// 	// 📌 **Önce Workers tablosunda arama yap**
// 	if err := config.DB.Where("email = ?", req.Email).First(&worker).Error; err == nil {
// 		// 🎯 **Worker bulundu, şifreyi kontrol et**
// 		if err := bcrypt.CompareHashAndPassword([]byte(worker.Password), []byte(req.Password)); err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre yanlış!"})
// 			return
// 		}

// 		// 📌 **JWT Token oluştur ve döndür**
// 		// token, _ := generateJWT(worker.Email, "worker")
// 		// c.JSON(http.StatusOK, gin.H{"token": token, "userType": "worker"})
// 		return
// 	}

// 	// 📌 **Worker bulunamadıysa, Customers tablosuna bak**
// 	if err := config.DB.Where("email = ?", req.Email).First(&customer).Error; err == nil {
// 		// 🎯 **Customer bulundu, şifreyi kontrol et**
// 		if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password)); err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre yanlış!"})
// 			return
// 		}

// 		// 📌 **JWT Token oluştur ve döndür**
// 		// token, _ := generateJWT(customer.Email, "customer")
// 		// c.JSON(http.StatusOK, gin.H{"token": token, "userType": "customer"})
// 		return
// 	}

// 	// 📌 **Eğer hiçbir kullanıcı bulunamazsa hata dön**
// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı!"})
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
	var contact models.ContactPerson

	// 📌 **Önce Workers tablosunda arama yap**
	if err := config.DB.Where("email = ?", req.Email).First(&worker).Error; err == nil {
		fmt.Println("✅ Worker bulundu:", worker.Email, "Şifre:", worker.Password) // Debug log

		// 🎯 **Şifreyi karşılaştır (Düz metin)**
		if worker.Password != req.Password {
			fmt.Println("❌ Girilen Şifre:", req.Password) // 🔥 Debug log
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre yanlış!"})
			return
		}

		// 📌 **JWT Token oluştur ve döndür**
		token := generateFakeJWT(worker.Email, "worker")
		c.JSON(http.StatusOK, gin.H{"token": token, "userType": "worker"})
		return
	}

	// 📌 **Worker bulunamadıysa, Contacts tablosuna bak**
	if err := config.DB.Where("email = ?", req.Email).First(&contact).Error; err == nil {
		fmt.Println("✅ Contact bulundu:", contact.Email, "Şifre:", contact.Password) // Debug log

		// 🎯 **Şifreyi karşılaştır (Düz metin)**
		if contact.Password != req.Password {
			fmt.Println("❌ Girilen Şifre:", req.Password) // 🔥 Debug log
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre yanlış!"})
			return
		}

		// 📌 **JWT Token oluştur ve döndür**
		token := generateFakeJWT(contact.Email, "customer")
		c.JSON(http.StatusOK, gin.H{"token": token, "userType": "contact"})
		return
	}

	// 📌 **Worker bulunamadıysa, Customers tablosuna bak**
	if err := config.DB.Where("email = ?", req.Email).First(&customer).Error; err == nil {
		fmt.Println("✅ Customer bulundu:", customer.Email, "Şifre:", customer.Password) // Debug log

		// 🎯 **Şifreyi karşılaştır (Düz metin)**
		if customer.Password != req.Password {
			fmt.Println("❌ Girilen Şifre:", req.Password) // 🔥 Debug log
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre yanlış!"})
			return
		}

		// 📌 **JWT Token oluştur ve döndür**
		token := generateFakeJWT(customer.Email, "customer")
		c.JSON(http.StatusOK, gin.H{"token": token, "userType": "customer"})
		return
	}

	// 📌 **Eğer hiçbir kullanıcı bulunamazsa hata dön**
	fmt.Println("❌ Kullanıcı bulunamadı:", req.Email) // 🔥 Debug log
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı!"})
}

func generateFakeJWT(email string, userType string) string {
	// ⚠ Gerçek JWT Kullanılmıyor! Sadece ID ve type içeren basit bir string döndürülüyor.
	return fmt.Sprintf("%s|%s|FAKE-TOKEN", email, userType)
}
