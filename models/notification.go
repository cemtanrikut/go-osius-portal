package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Message   string    `json:"message"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}
