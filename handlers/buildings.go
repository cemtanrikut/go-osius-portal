package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/config"
	"main.go/models"
)

// **Yeni ID'yi oluşturma fonksiyonu**
func generateBuildingID(db *gorm.DB) (string, error) {
	fmt.Println("📌 Yeni bir bina ekleme isteği alındı!") // 🛠 Log ekleyelim

	var lastBuilding models.Building

	// **En son eklenen binayı ID'ye göre sırala ve getir**
	if err := db.Order("id DESC").First(&lastBuilding).Error; err != nil {
		// **Eğer hiç bina yoksa, "B-0001" ile başlat**
		if err == gorm.ErrRecordNotFound {
			return "B-0001", nil
		}
		return "", err
	}

	// **Mevcut en büyük ID'yi al ("B-xxxx" formatında)**
	lastID := lastBuilding.ID

	// **"B-" kısmını at ve sayısal kısmı çıkar**
	lastNumberStr := strings.TrimPrefix(lastID, "B-")
	lastNumber, err := strconv.Atoi(lastNumberStr)
	if err != nil {
		return "", fmt.Errorf("invalid ID format: %v", err)
	}

	// **Bir artırarak yeni ID'yi oluştur**
	newID := fmt.Sprintf("B-%04d", lastNumber+1)
	return newID, nil
}

// 📌 **Yeni Bina Ekleme**
func CreateBuilding(c *gin.Context) {
	fmt.Println("📌 Yeni bir bina ekleme isteği alındı! 2222") // 🛠 Log ekleyelim

	var building models.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 📌 **Yeni ID oluştur**
	newID, err := generateBuildingID(config.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate building ID"})
		return
	}
	building.ID = newID

	// 📌 **Veritabanına ekle**
	if err := config.DB.Create(&building).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create building"})
		return
	}

	// 📌 **Başarı yanıtı**
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Building added successfully",
		"building": building,
	})
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
