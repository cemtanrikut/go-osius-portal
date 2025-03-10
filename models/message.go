package models

import "gorm.io/gorm"

// 📩 Message Modeli (Mesajlar)
type Message struct {
	gorm.Model
	TicketID uint   `json:"ticket_id"` // 🎫 Hangi Ticket'a ait?
	Sender   string `json:"sender"`    // 📩 Gönderen (You / Diğer kullanıcı)
	Text     string `json:"text"`      // 📄 Mesaj İçeriği
	FileURL  string `json:"file_url"`  // 📂 Eğer dosya varsa
	FileType string `json:"file_type"` // 🖼 image / file
	Time     string `json:"time"`      // ⏰ Gönderim zamanı
}
