package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// 📌 **Kullanıcı Girişi (Login)**
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 📌 **Şifre kontrolü yapılmalı! (Burada basit bir string karşılaştırması var, hashleme eklenebilir)**
	if user.Password != existingUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": existingUser})
}

// 📌 **Kullanıcı Çıkışı (Logout)**
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
