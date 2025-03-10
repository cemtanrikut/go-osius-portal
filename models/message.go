package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	TicketID uint   `json:"ticket_id"`
	Sender   string `json:"sender"`
	Text     string `json:"text"`
	FileURL  string `json:"file_url,omitempty"` // If file exists will add url
}
