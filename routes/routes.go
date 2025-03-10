package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

// SetupRouter - Tüm route'ları burada tanımlıyoruz
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 📌 **Ticket İşlemleri**
	ticketRoutes := r.Group("/tickets")
	{
		ticketRoutes.GET("", handlers.GetTickets)          // Tüm ticket'ları getir
		ticketRoutes.GET("/:id", handlers.GetTicketByID)   // Belirli ticket'ı getir
		ticketRoutes.POST("", handlers.CreateTicket)       // Yeni ticket oluştur
		ticketRoutes.PUT("/:id", handlers.UpdateTicket)    // Ticket güncelle
		ticketRoutes.DELETE("/:id", handlers.DeleteTicket) // Ticket sil
	}

	// 📌 **Mesaj İşlemleri (Belirli Ticket İçin)**
	messageRoutes := r.Group("/tickets/:id/messages")
	{
		messageRoutes.GET("", handlers.GetMessages)    // Ticket'a ait mesajları getir
		messageRoutes.POST("", handlers.CreateMessage) // Ticket'a yeni mesaj ekle
	}

	// 📌 **Dosya Yükleme (Resim/Dosya)**
	r.POST("/upload", handlers.UploadFile) // Dosya yükleme endpoint'i

	// 📌 **Bina İşlemleri**
	buildingRoutes := r.Group("/buildings")
	{
		buildingRoutes.GET("", handlers.GetBuildings)          // Tüm binaları getir
		buildingRoutes.GET("/:id", handlers.GetBuildingByID)   // Belirli bir binayı getir
		buildingRoutes.POST("", handlers.CreateBuilding)       // Yeni bina ekle
		buildingRoutes.PUT("/:id", handlers.UpdateBuilding)    // Binayı güncelle
		buildingRoutes.DELETE("/:id", handlers.DeleteBuilding) // Binayı sil
	}

	return r
}
