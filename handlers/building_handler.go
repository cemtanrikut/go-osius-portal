package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// ✅ Creates Building (POST /buildings)
func CreateBuilding(c *gin.Context) {
	var building models.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&building)
	c.JSON(http.StatusCreated, building)
}

// ✅ Gets all buildings (GET /buildings)
func GetBuildings(c *gin.Context) {
	var buildings []models.Building
	config.DB.Find(&buildings)
	c.JSON(http.StatusOK, buildings)
}

// ✅ Gets building by id (GET /buildings/:id)
func GetBuildingByID(c *gin.Context) {
	var building models.Building
	id := c.Param("id")

	if err := config.DB.First(&building, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}

	c.JSON(http.StatusOK, building)
}

// ✅ Update building (PUT /buildings/:id)
func UpdateBuilding(c *gin.Context) {
	var building models.Building
	id := c.Param("id")

	if err := config.DB.First(&building, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}

	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&building)
	c.JSON(http.StatusOK, building)
}

// ✅ Delete building (DELETE /buildings/:id)
func DeleteBuilding(c *gin.Context) {
	var building models.Building
	id := c.Param("id")

	if err := config.DB.First(&building, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}

	config.DB.Delete(&building)
	c.JSON(http.StatusOK, gin.H{"message": "Building deleted succesfully!"})
}
