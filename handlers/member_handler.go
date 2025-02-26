package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// ✅ Creates member (POST /members)
func CreateMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&member)
	c.JSON(http.StatusCreated, member)
}

// ✅ Gets all members (GET /members)
func GetMembers(c *gin.Context) {
	var members []models.Member
	config.DB.Find(&members)
	c.JSON(http.StatusOK, members)
}

// ✅ Gets member by id (GET /members/:id)
func GetMemberByID(c *gin.Context) {
	var member models.Member
	id := c.Param("id")

	if err := config.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	c.JSON(http.StatusOK, member)
}

// ✅ Updates a member (PUT /members/:id)
func UpdateMember(c *gin.Context) {
	var member models.Member
	id := c.Param("id")

	if err := config.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&member)
	c.JSON(http.StatusOK, member)
}

// ✅ Deletes a member (DELETE /members/:id)
func DeleteMember(c *gin.Context) {
	var member models.Member
	id := c.Param("id")

	if err := config.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	config.DB.Delete(&member)
	c.JSON(http.StatusOK, gin.H{"message": "Member deleted succesfully!"})
}
