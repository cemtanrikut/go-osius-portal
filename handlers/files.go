package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"main.go/config"
	"main.go/models"

	"github.com/gin-gonic/gin"
)

// 📌 **Dosya Yükleme ve Ticket'a Bağlama**
func UploadFile(c *gin.Context) {
	ticketID := c.Param("ticketId")

	// Ticket ID'nin var olup olmadığını kontrol et
	var ticket models.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	// Çoklu dosya yükleme desteği
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload request"})
		return
	}

	files := form.File["files"]
	var savedFiles []models.File

	for _, file := range files {
		// Dosya adını belirle
		filePath := fmt.Sprintf("uploads/%s", file.Filename)

		// Dosyayı sunucuya kaydet
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// Dosya bilgilerini oluştur
		fileRecord := models.File{
			TicketID: ticketID,
			Filename: file.Filename,
			FileURL:  filePath, // Local dosya yolu
			FileType: filepath.Ext(file.Filename),
		}

		// Veritabanına kaydet
		config.DB.Create(&fileRecord)
		savedFiles = append(savedFiles, fileRecord)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully", "files": savedFiles})
}

// 📌 **Bir Ticket'a Ait Dosyaları Getirme**
func GetFilesByTicketID(c *gin.Context) {
	ticketID := c.Param("ticketId")
	var files []models.File
	config.DB.Where("ticket_id = ?", ticketID).Find(&files)
	c.JSON(http.StatusOK, files)
}

// 📌 **Dosya Silme**
func DeleteFile(c *gin.Context) {
	fileID := c.Param("fileId")
	var file models.File
	if err := config.DB.First(&file, fileID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Dosyayı veritabanından sil
	config.DB.Delete(&file)
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
