package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"main.go/models"

	"main.go/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 📌 **WebSocket Upgrader (HTTP'yi WebSocket'e dönüştürmek için)**
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true, // 🚀 **TLS desteği için ekledik**
}

// 📌 **WebSocket Bağlantılarını Yönetmek İçin Map**
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)
var mutex = sync.Mutex{} // Çoklu işlem için senkronizasyon

// 📌 **WebSocket Bağlantısını Açma**
func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("❌ WebSocket bağlantı hatası:", err)
		return
	}

	// 📌 **Yeni bağlantıyı kaydet**
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	fmt.Println("✅ Yeni WebSocket bağlantısı:", conn.RemoteAddr())

	// 📌 **Bağlantıyı dinle**
	go func() {
		defer func() {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			conn.Close()
			fmt.Println("❌ WebSocket bağlantısı kapandı:", conn.RemoteAddr())
		}()

		for {
			var msg models.Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Println("❌ Mesaj okuma hatası:", err)
				break
			}

			// 📌 **Mesajı veritabanına kaydet**
			if err := config.DB.Create(&msg).Error; err != nil {
				fmt.Println("❌ DB'ye mesaj kaydedilemedi:", err)
				continue
			}

			// 📌 **Mesajı yayına gönder**
			broadcast <- msg
		}
	}()
}

// 📌 **WebSocket Üzerinden Mesajları Yayınlama**
func BroadcastMessages() {
	for {
		msg := <-broadcast

		mutex.Lock()
		for conn := range clients {
			err := conn.WriteJSON(msg)
			if err != nil {
				fmt.Println("❌ Mesaj gönderme hatası:", err)
				conn.Close()
				delete(clients, conn)
			} else {
				fmt.Println("📩 Mesaj gönderildi:", msg.Text)
			}
		}
		mutex.Unlock()
	}
}

func GetMessages(c *gin.Context) {
	ticketID := c.Param("ticketId")

	var messages []models.Message
	if err := config.DB.Where("ticket_id = ?", ticketID).Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Mesajlar yüklenemedi"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func CreateMessage(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 📌 **Mesajı DB'ye kaydet**
	if err := config.DB.Create(&msg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Mesaj kaydedilemedi"})
		return
	}

	// 📌 **WebSocket Yayınına Gönder**
	broadcast <- msg

	c.JSON(http.StatusOK, msg)
}

func DeleteMessage(c *gin.Context) {
	messageID := c.Param("messageId")

	if err := config.DB.Delete(&models.Message{}, messageID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Mesaj silinemedi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mesaj başarıyla silindi"})
}
