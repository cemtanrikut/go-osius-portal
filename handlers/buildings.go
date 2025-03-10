package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// 📌 **Yeni Bina Ekleme**
func CreateBuilding(c *gin.Context) {
	var building models.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&building)
	c.JSON(http.StatusCreated, building)
}

// 📌 **Tüm Binaları Getirme**
func GetBuildings(c *gin.Context) {
	var buildings []models.Building
	config.DB.Find(&buildings)
	c.JSON(http.StatusOK, buildings)
}

// 📌 **Belirli Bir Binayı Getirme**
func GetBuildingByID(c *gin.Context) {
	var building models.Building
	id := c.Param("id")

	if err := config.DB.First(&building, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}

	c.JSON(http.StatusOK, building)
}

// 📌 **Bina Güncelleme**
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

// 📌 **Bina Silme**
func DeleteBuilding(c *gin.Context) {
	var building models.Building
	id := c.Param("id")

	if err := config.DB.First(&building, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}

	config.DB.Delete(&building)
	c.JSON(http.StatusOK, gin.H{"message": "Building deleted successfully"})
}
