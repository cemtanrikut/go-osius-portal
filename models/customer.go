package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
