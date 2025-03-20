package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	ID                   string          `json:"id"`
	Name                 string          `json:"name"`
	Address              string          `json:"address"`
	PostCode             string          `json:"postCode"`
	Plaats               string          `json:"plaats"`
	Country              string          `json:"country"`
	Phone                string          `json:"phone"`
	Email                string          `json:"email"`
	Password             string          `json:"password"`
	Website              string          `json:"website"`
	Logo                 *string         `json:"logo"`
	Status               string          `json:"status"`
	Supplier             bool            `json:"supplier"`
	BtwNumber            string          `json:"btwNumber"`
	KvK                  string          `json:"kvk"`
	Vestigingsnummer     string          `json:"vestigingsnummer"`
	Relatiebeheerder     string          `json:"relatiebeheerder"`
	GlobalLocationNumber string          `json:"globalLocationNumber"`
	Moederonderneming    string          `json:"moederonderneming"`
	Remarks              *string         `json:"remarks"`
	Contacts             []ContactPerson `json:"contacts" gorm:"-"`
}

type ContactPerson struct {
	gorm.Model
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	BuildingID string `json:"buildingId"`
	Role       string `json:"role"`
	FirstName  string `json:"firstName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
}
