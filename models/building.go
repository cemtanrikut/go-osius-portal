package models

import "gorm.io/gorm"

type Building struct {
	gorm.Model
	ID            string `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	HouseNo       string `json:"houseNo"`
	PostCode      string `json:"postCode"`
	Plaats        string `json:"plaats"`
	Status        string `json:"status"`
	Note          string `json:"note"`
	CalculateType string `json:"calculateType"`
}
