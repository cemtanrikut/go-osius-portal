package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// CreateTicket godoc
// @Summary Creates a new ticket
// @Description Adds a new ticket to the database
// @Tags tickets
// @Accept  json
// @Produce  json
// @Param   ticket  body     models.Ticket  true  "Ticket data"
// @Success 201      {object} models.Ticket
// @Failure 400      {object} map[string]string
// @Router  /tickets [post]
func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&ticket)
	c.JSON(http.StatusCreated, ticket)
}

// GetTickets godoc
// @Summary Retrieves all tickets
// @Description Gets a list of all tickets
// @Tags tickets
// @Produce  json
// @Success 200  {array}  models.Ticket
// @Failure 500  {object} map[string]string
// @Router  /tickets [get]
func GetTickets(c *gin.Context) {
	var tickets []models.Ticket
	config.DB.Find(&tickets)
	c.JSON(http.StatusOK, tickets)
}

// GetTicketByID godoc
// @Summary Retrieves a ticket by ID
// @Description Gets a specific ticket using its ID
// @Tags tickets
// @Produce  json
// @Param   id   path     int  true  "Ticket ID"
// @Success 200  {object} models.Ticket
// @Failure 404  {object} map[string]string
// @Router  /tickets/{id} [get]
func GetTicketByID(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")

	if err := config.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// UpdateTicket godoc
// @Summary Updates a ticket
// @Description Updates an existing ticket by ID
// @Tags tickets
// @Accept  json
// @Produce  json
// @Param   id        path     int              true  "Ticket ID"
// @Param   ticket    body     models.Ticket  true  "Updated ticket data"
// @Success 200      {object} models.Ticket
// @Failure 400      {object} map[string]string
// @Failure 404      {object} map[string]string
// @Router  /tickets/{id} [put]
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

// DeleteTicket godoc
// @Summary Deletes a ticket
// @Description Deletes a specific ticket using its ID
// @Tags tickets
// @Param   id   path     int  true  "Ticket ID"
// @Success 200  {object} map[string]string
// @Failure 404  {object} map[string]string
// @Router  /tickets/{id} [delete]
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
