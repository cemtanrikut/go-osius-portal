package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

func generateCustomerID() (string, error) {
	var lastCustomer models.Customer
	var lastID int

	if err := config.DB.Order("id DESC").First(&lastCustomer).Error; err == nil {
		// "C-0001" formatından sayıyı çekiyoruz
		fmt.Sscanf(lastCustomer.ID, "C-%d", &lastID)
	}

	newID := fmt.Sprintf("C-%04d", lastID+1)
	return newID, nil
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Yeni ID üret
	newID, err := generateCustomerID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate customer ID"})
		return
	}

	customer.ID = newID // Yeni ID'yi ata

	// DB'ye kaydet
	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

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

	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// 📌 **Müşteri Güncelleme**
func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
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

	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	config.DB.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
