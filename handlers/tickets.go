package handlers

import (
	"fmt"
	"net/http"

	"main.go/config"
	"main.go/models"

	"github.com/gin-gonic/gin"
)

// 📌 **Son Ticket ID'yi alıp yeni bir ID oluşturma fonksiyonu**
func generateTicketID() string {
	var lastTicket models.Ticket
	config.DB.Order("id DESC").First(&lastTicket)

	if lastTicket.TicketId == "" {
		return "T-0001"
	}

	// Son ID'yi al, integer kısmını ayır, artır ve tekrar formatla
	var lastIDNumber int
	fmt.Sscanf(fmt.Sprintf("T-%04d", lastTicket.ID), "T-%04d", &lastIDNumber)
	newID := fmt.Sprintf("T-%04d", lastIDNumber+1)
	return newID
}

// 📌 **Yeni Ticket Ekleme**
func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Yeni Ticket ID oluştur
	ticket.TicketId = generateTicketID()

	// Veritabanına kaydet
	config.DB.Create(&ticket)
	c.JSON(http.StatusCreated, ticket)
}

// 📌 **Tüm Ticketları Listeleme**
func GetTickets(c *gin.Context) {
	var tickets []models.Ticket
	result := config.DB.Unscoped().Preload("Files").Find(&tickets) // 🔥 Silinmiş kayıtları da getir

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	fmt.Println("📌 API'den Dönen Tüm Ticketlar:", tickets) // ✅ Terminalde kontrol et
	c.JSON(http.StatusOK, tickets)
}

// 📌 **Tek Bir Ticket Getirme**
func GetTicketByID(c *gin.Context) {
	id := c.Param("id")
	var ticket models.Ticket
	if err := config.DB.Preload("Files").First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// 📌 **Ticket Güncelleme**
func UpdateTicket(c *gin.Context) {
	id := c.Param("id")                  // URL'den gelen ID
	fmt.Println("🛠 Gelen ticketId:", id) // 🔥 Debug için ekrana basalım

	var ticket models.Ticket
	if err := config.DB.First(&ticket, "ticket_id = ?", id).Error; err != nil {
		fmt.Println("❌ Ticket bulunamadı, hata:", err) // 🔥 Hata mesajını ekrana basalım
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	// JSON'dan gelen verileri ekrana basalım
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		fmt.Println("❌ JSON Bind Hatası:", err) // JSON formatı yanlış olabilir
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("📝 Güncellenecek Veri:", requestData) // 🔥 Backend'e gelen datayı kontrol et!

	// Ticket güncelleniyor
	if err := config.DB.Model(&ticket).Updates(requestData).Error; err != nil {
		fmt.Println("❌ Güncelleme hatası:", err) // Eğer güncelleme başarısız olursa
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	fmt.Println("✅ Güncelleme başarılı:", ticket) // Güncellenmiş ticket'ı göster
	c.JSON(http.StatusOK, ticket)
}

// 📌 **Ticket Silme**
func DeleteTicket(c *gin.Context) {
	id := c.Param("id")
	var ticket models.Ticket
	if err := config.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	// Önce ilgili dosyaları sil
	config.DB.Where("ticket_id = ?", id).Delete(&models.File{})

	// Sonra ticket'ı sil
	config.DB.Delete(&ticket)
	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
