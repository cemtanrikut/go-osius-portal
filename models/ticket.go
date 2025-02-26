package models

type Ticket struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `bson:"title" json:"title"`
	Company     string `bson:"company" json:"company"`
	Location    string `bson:"location" json:"location"`
	ReferanceNo string `bson:"referance_no" json:"referance_no"`
	Person      string `bson:"person" json:"person"`
	TicketType  string `bson:"ticket_type" json:"ticket_type"`
	Description string `bson:"description" json:"description"`
	Files       string `bson:"files" json:"files"`
	Status      string `bson:"status" json:"status"`
	CreateDate  string `bson:"create_date" json:"create_date"`
}
