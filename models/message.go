package models

// 📌 **Mesaj Modeli**
type Message struct {
	TicketID uint   `json:"ticket_id" gorm:"index"`
	Sender   string `json:"sender"`
	Text     string `json:"text"`
	FileURL  string `json:"file_url,omitempty"` // Eğer varsa
}
