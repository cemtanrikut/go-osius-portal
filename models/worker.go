package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	ID         string `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Role       string `json:"role"`
	StartDate  string `json:"startDate"`
	Status     string `json:"status"`
}
