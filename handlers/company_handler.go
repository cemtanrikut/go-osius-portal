package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

func CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&company)
	c.JSON(http.StatusOK, company)
}

func GetCompanies(c *gin.Context) {
	var companies []models.Company
	config.DB.Find(&companies)
	c.JSON(http.StatusOK, companies)
}

func GetCompanyByID(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

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
