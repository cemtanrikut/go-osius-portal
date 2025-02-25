package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// ✅ Creates company (POST /companies)
func CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&company)
	c.JSON(http.StatusOK, company)
}

// ✅ Gets all companies (GET /companies)
func GetCompanies(c *gin.Context) {
	var companies []models.Company
	config.DB.Find(&companies)
	c.JSON(http.StatusOK, companies)
}

// ✅ GetCompanyByID (GET /companies/:id)
func GetCompanyByID(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

// ✅ Update company (PUT /companies/:id)
func UpdateCompany(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&company)
	c.JSON(http.StatusOK, company)
}

// ✅ Delete company (DELETE /companies/:id)
func DeleteCompany(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	config.DB.Delete(&company)
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted succesfully!"})
}
