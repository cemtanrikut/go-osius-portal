package models

type Company struct {
	ID                    uint                  `gorm:"primaryKey"`
	Name                  string                `bson:"name" json:"name"`
	Phone                 string                `bson:"phone" json:"phone"`
	Email                 string                `bson:"email" json:"email"`
	Tags                  string                `bson:"tags" json:"tags"`
	Status                bool                  `bson:"status" json:"status"`
	CreatedAt             string                `bson:"created_at" json:"created_at"`
	CompanyInfo           CompanyInfo           `bson:"company_info" json:"company_info" gorm:"embedded"`
	CompanyRepresentative CompanyRepresentative `bson:"company_representative" json:"company_representative" gorm:"embedded"`
}

type CompanyInfo struct {
	CompanyName string `bson:"company_name" json:"company_name"`
	Address     string `bson:"address" json:"address"`
	Postcode    string `bson:"postcode" json:"postcode"`
	Country     string `bson:"country" json:"country"`
	KVK         string `bson:"kvk" json:"kvk"`
	BTWnumber   string `bson:"btw" json:"btw"`
	LastUpdate  string `bson:"last_update" json:"last_update"`
}

type CompanyRepresentative struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Phone string `bson:"phone" json:"phone"`
}
