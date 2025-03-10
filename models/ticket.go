package models

import "gorm.io/gorm"

// 🎫 Ticket Modeli
type Ticket struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	AssignedTo  string `json:"assignedTo"`
	Date        string `json:"date"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	CreatedBy   string `json:"createdBy"`
}
