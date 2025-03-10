package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// 🎯 1️⃣ Belirli Bir Ticket'ın Mesajlarını Getir (GET /tickets/:id/messages)
func GetMessages(c *gin.Context) {
	var messages []models.Message
	ticketID := c.Param("id")

	if err := config.DB.Where("ticket_id = ?", ticketID).Find(&messages).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Messages not found"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// 🎯 2️⃣ Belirli Bir Ticket'a Mesaj Ekle (POST /tickets/:id/messages)
func CreateMessage(c *gin.Context) {
	var message models.Message
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.TicketID = uint(ticketID) // 🎫 Ticket ile ilişkilendirme
	config.DB.Create(&message)
	c.JSON(http.StatusCreated, message)
}
