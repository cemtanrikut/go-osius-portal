package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// ✅ Creates Ticket (POST /tickets)
func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&ticket)
	c.JSON(http.StatusCreated, ticket)
}

// ✅ Gets all tickets (GET /tickets)
func GetTickets(c *gin.Context) {
	var tickets []models.Ticket
	config.DB.Find(&tickets)
	c.JSON(http.StatusOK, tickets)
}

// ✅ Gets a ticket by id (GET /tickets/:id)
func GetTicketByID(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")

	if err := config.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// ✅ Update a ticket by id (PUT /tickets/:id)
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

// ✅ Deletes a ticket by id (DELETE /tickets/:id)
func DeleteTicket(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")

	if err := config.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	config.DB.Delete(&ticket)
	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted succesfully"})
}
