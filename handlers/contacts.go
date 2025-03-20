package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/models"
)

func generateContactID() (string, error) {
	var lastContact models.ContactPerson
	var lastID int

	if err := config.DB.Order("id DESC").First(&lastContact).Error; err == nil {
		// "C-0001" formatından sayıyı çekiyoruz
		fmt.Sscanf(lastContact.ID, "CP-%d", &lastID)
	}

	newID := fmt.Sprintf("CP-%04d", lastID+1)
	return newID, nil
}

func CreateContact(c *gin.Context) {
	var contact models.ContactPerson

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Yeni ID üret
	newID, err := generateContactID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate contact ID"})
		return
	}

	contact.ID = newID // Yeni ID'yi ata

	// DB'ye kaydet
	if err := config.DB.Create(&contact).Error; err != nil {
		log.Println("Error inserting contact:", err) // 🔹 Daha fazla hata detayı logla
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contact"})
		return
	}

	c.JSON(http.StatusCreated, contact)
}

// 📌 **Tüm Customerlarin Contact Personlari Getirme**
func GetCustomerContacts(c *gin.Context) {
	var contacts []models.ContactPerson
	config.DB.Find(&contacts)
	c.JSON(http.StatusOK, contacts)
}

// 📌 **Tüm Buildinglerin Contact Personlari Getirme**
func GetBuildingContacts(c *gin.Context) {
	var contacts []models.ContactPerson
	config.DB.Find(&contacts)
	c.JSON(http.StatusOK, contacts)
}

// 📌 **Belirli Bir Contact Person Getirme**
func GetCustomerContactByID(c *gin.Context) {
	var contacts []models.ContactPerson // DİKKAT: Liste olarak tanımladık
	id := c.Param("customerId")

	if err := config.DB.Where("customer_id = ?", id).Find(&contacts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer's contacts not found"})
		return
	}

	c.JSON(http.StatusOK, contacts) // Tekil obje yerine ARRAY döndür
}

// 📌 **Belirli Bir Contact Person Getirme**
func GetBuildingContactByID(c *gin.Context) {
	var contacts []models.ContactPerson // DİKKAT: Liste olarak tanımladık
	id := c.Param("buildingId")

	if err := config.DB.Where("building_id = ?", id).Find(&contacts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buildings's contacts not found"})
		return
	}

	c.JSON(http.StatusOK, contacts) // Tekil obje yerine ARRAY döndür
}

// 📌 **Müşteri Contact Person Güncelleme**
func UpdateContact(c *gin.Context) {
	var contact models.ContactPerson
	id := c.Param("id")

	if err := config.DB.First(&contact, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&contact)
	c.JSON(http.StatusOK, contact)
}

// 📌 **Contact Silme**
func DeleteContact(c *gin.Context) {
	var contact models.ContactPerson
	id := c.Param("id")

	if err := config.DB.First(&contact, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	config.DB.Delete(&contact)
	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}
