package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// 🎯 1️⃣ Tüm Ticket'ları Listele (GET /tickets)
func GetTickets(c *gin.Context) {
	var tickets []models.Ticket
	config.DB.Find(&tickets)
	c.JSON(http.StatusOK, tickets)
}

// 🎯 2️⃣ Yeni Ticket Oluştur (POST /tickets)
func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&ticket)
	c.JSON(http.StatusCreated, ticket)
}

// 🎯 3️⃣ Belirli Bir Ticket'ı Getir (GET /tickets/:id)
func GetTicketByID(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")

	if err := config.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// 🎯 4️⃣ Ticket Güncelle (PUT /tickets/:id)
func UpdateTicket(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")

	if err := config.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&ticket)
	c.JSON(http.StatusOK, ticket)
}

// 🎯 5️⃣ Ticket Sil (DELETE /tickets/:id)
func DeleteTicket(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")

	if err := config.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	config.DB.Delete(&ticket)
	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
