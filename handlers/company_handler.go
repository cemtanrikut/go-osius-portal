package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// CreateCompany godoc
// @Summary Creates a new company
// @Description Adds a new company to the database
// @Tags companies
// @Accept  json
// @Produce  json
// @Param   company  body     models.Company  true  "Company data"
// @Success 201      {object} models.Company
// @Failure 400      {object} map[string]string
// @Router  /companies [post]
func CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&company)
	c.JSON(http.StatusOK, company)
}

// GetCompanies godoc
// @Summary Retrieves all companies
// @Description Gets a list of all companies
// @Tags companies
// @Produce  json
// @Success 200  {array}  models.Company
// @Failure 500  {object} map[string]string
// @Router  /companies [get]
func GetCompanies(c *gin.Context) {
	var companies []models.Company
	config.DB.Find(&companies)
	c.JSON(http.StatusOK, companies)
}

// GetCompanyByID godoc
// @Summary Retrieves a company by ID
// @Description Gets a specific company using its ID
// @Tags companies
// @Produce  json
// @Param   id   path     int  true  "Company ID"
// @Success 200  {object} models.Company
// @Failure 404  {object} map[string]string
// @Router  /companies/{id} [get]
func GetCompanyByID(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

// @Summary Updates a company
// @Description Updates an existing company by ID
// @Tags companies
// @Accept  json
// @Produce  json
// @Param   id        path     int              true  "Company ID"
// @Param   company  body     models.Company  true  "Updated company data"
// @Success 200      {object} models.Company
// @Failure 400      {object} map[string]string
// @Failure 404      {object} map[string]string
// @Router  /companies/{id} [put]
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

// DeleteCompany godoc
// @Summary Deletes a company
// @Description Deletes a specific company using its ID
// @Tags companies
// @Param   id   path     int  true  "Company ID"
// @Success 200  {object} map[string]string
// @Failure 404  {object} map[string]string
// @Router  /companies/{id} [delete]
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
