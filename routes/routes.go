package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

// SetupRouter - Tüm route'ları burada tanımlıyoruz
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 🔥 CORS middleware ekleyelim
	r.Use(cors.Default())

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

	// 📌 **Müşteri İşlemleri**
	customerRoutes := r.Group("/customers")
	{
		customerRoutes.GET("", handlers.GetCustomers)          // Tüm müşterileri getir
		customerRoutes.GET("/:id", handlers.GetCustomerByID)   // Belirli bir müşteriyi getir
		customerRoutes.POST("", handlers.CreateCustomer)       // Yeni müşteri ekle
		customerRoutes.PUT("/:id", handlers.UpdateCustomer)    // Müşteriyi güncelle
		customerRoutes.DELETE("/:id", handlers.DeleteCustomer) // Müşteriyi sil
	}

	workerRoutes := r.Group("/workers")
	{
		workerRoutes.GET("", handlers.GetWorkers)            // Tüm binaları getir
		workerRoutes.GET("/:id", handlers.GetBuildingByID)   // Belirli bir binayı getir
		workerRoutes.POST("", handlers.CreateWorker)         // Yeni bina ekle
		workerRoutes.PUT("/:id", handlers.UpdateBuilding)    // Binayı güncelle
		workerRoutes.DELETE("/:id", handlers.DeleteBuilding) // Binayı sil
	}

	// 📌 **Auth İşlemleri**
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", handlers.Login)   // Kullanıcı giriş yapar
		authRoutes.POST("/logout", handlers.Logout) // Kullanıcı çıkış yapar
	}

	// 📌 **Bildirim İşlemleri**
	notificationRoutes := r.Group("/notifications")
	{
		notificationRoutes.GET("", handlers.GetNotifications)                // Tüm bildirimleri getir
		notificationRoutes.PUT("/:id/read", handlers.MarkNotificationAsRead) // Bildirimi okundu olarak işaretle
	}

	// 📌 **Dashboard Verileri**
	r.GET("/dashboard", handlers.GetDashboardData) // 📊 Dashboard verilerini getir

	return r
}
