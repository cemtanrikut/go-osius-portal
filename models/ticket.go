package models

import (
	"gorm.io/gorm"
)

// Ticket Modeli
type Ticket struct {
	gorm.Model
	TicketId         string `json:"ticketId"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	WorkerID         string `json:"workerId"` // Foreign Key
	Worker           string `gorm:"foreignKey:WorkerID"`
	BuildingID       string `json:"buildingId"`
	Building         string `gorm:"foreignKey:BuildingID"`
	CustomerID       string `json:"customerId"`
	Customer         string `gorm:"foreignKey:CustomerID"`
	NotificationType string `json:"notificationType"`
	Date             string `json:"date"` // Tarih formatı düzeltilmiş
	Files            []File `gorm:"foreignKey:TicketID"`
	Status           string `json:"status"` // ToDo, InProgress, Done
	CreatorID        string `json:"creatorId"`
	CreatedBy        string `json:"created_by"` // Ticket'i oluşturan kullanıcı
}

// File Modeli
type File struct {
	gorm.Model
	TicketID string `json:"ticketId"` // Ticket ile ilişkilendirme
	Filename string `json:"filename"` // Orijinal dosya adı
	FileURL  string `json:"fileUrl"`  // Dosyanın saklandığı URL (Cloud veya Local)
	FileType string `json:"fileType"` // image/png, application/pdf, vb.
}
