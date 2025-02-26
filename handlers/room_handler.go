package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// ✅ Creates room (POST /rooms)
func CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&room)
	c.JSON(http.StatusCreated, room)
}

// ✅ Gets all rooms (GET /rooms)
func GetRooms(c *gin.Context) {
	var rooms []models.Room
	config.DB.Find(&rooms)
	c.JSON(http.StatusOK, rooms)
}

// ✅ Gets room by id (GET /rooms/:id)
func GetRoomByID(c *gin.Context) {
	var room models.Room
	id := c.Param("id")

	if err := config.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	c.JSON(http.StatusOK, room)
}

// ✅ Updates room (PUT /rooms/:id)
func UpdateRoom(c *gin.Context) {
	var room models.Room
	id := c.Param("id")

	if err := config.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&room)
	c.JSON(http.StatusOK, room)
}

// ✅ Deletes room by id (DELETE /rooms/:id)
func DeleteRoom(c *gin.Context) {
	var room models.Room
	id := c.Param("id")

	if err := config.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	config.DB.Delete(&room)
	c.JSON(http.StatusOK, gin.H{"message": "Room deleted succesfull!"})
}
