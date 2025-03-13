package handlers

import (
	"fmt"
	"net/http"

	"main.go/config"
	"main.go/models"

	"github.com/gin-gonic/gin"
)

// 📌 **Son ID'yi alıp bir sonraki ID'yi oluşturma fonksiyonu**
func generateWorkerID() string {
	var lastWorker models.Worker
	config.DB.Order("id DESC").First(&lastWorker)

	if lastWorker.ID == "" {
		return "W-0001"
	}

	// Son ID'yi al, integer kısmını ayır, artır ve tekrar formatla
	var lastIDNumber int
	fmt.Sscanf(lastWorker.ID, "W-%04d", &lastIDNumber)
	newID := fmt.Sprintf("W-%04d", lastIDNumber+1)
	return newID
}

// 📌 **Yeni Worker Ekleme**
func CreateWorker(c *gin.Context) {
	var worker models.Worker
	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Yeni Worker ID oluştur
	worker.ID = generateWorkerID()

	config.DB.Create(&worker)
	c.JSON(http.StatusCreated, worker)
}

// 📌 **Tüm Workers Listeleme**
func GetWorkers(c *gin.Context) {
	var workers []models.Worker
	config.DB.Find(&workers)
	c.JSON(http.StatusOK, workers)
}
