package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

// 📌 **Yeni Müşteri Ekleme**
func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&customer)
	c.JSON(http.StatusCreated, customer)
}

// 📌 **Tüm Müşterileri Getirme**
func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	config.DB.Find(&customers)
	c.JSON(http.StatusOK, customers)
}

// 📌 **Belirli Bir Müşteriyi Getirme**
func GetCustomerByID(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// 📌 **Müşteri Güncelleme**
func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&customer)
	c.JSON(http.StatusOK, customer)
}

// 📌 **Müşteri Silme**
func DeleteCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	config.DB.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
