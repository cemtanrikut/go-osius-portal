package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// 📌 **Tüm Bildirimleri Getirme**
func GetNotifications(c *gin.Context) {
	var notifications []models.Notification
	config.DB.Order("created_at desc").Find(&notifications)
	c.JSON(http.StatusOK, notifications)
}

// 📌 **Bildirim Okundu Olarak İşaretleme**
func MarkNotificationAsRead(c *gin.Context) {
	id := c.Param("id")
	var notification models.Notification

	if err := config.DB.First(&notification, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	notification.Read = true
	config.DB.Save(&notification)
	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}
