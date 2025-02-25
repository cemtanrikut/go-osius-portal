package models

type Building struct {
	ID              string `bson:"_id,omitempty" json:"id"`
	LocationName    string `bson:"name" json:"name"`
	LocationAddress string `bson:"address" json:"address"`
	HouseNumber     string `bson:"house_number" json:"house_number"`
	PostalCode      string `bson:"postal_code" json:"postal_code"`
	Plaats          string `bson:"plaats" json:"plaats"`
	Email           string `bson:"email" json:"email"`
	Status          bool   `bson:"status" json:"status"`
	Note            string `bson:"note" json:"note"`
	CalculateType   string `bson:"calculate_type" json:"calculate_type"`
	CreatedDate     string `bson:"created_date" json:"created_date"`
}

type Member struct {
	ID                        string `bson:"_id,omitempty" json:"id"`
	BuildingID                string `bson:"building_id" json:"building_id"`
	Name                      string `bson:"name" json:"name"`
	Email                     string `bson:"email" json:"email"`
	Phone                     string `bson:"phone" json:"phone"`
	Postcode                  string `bson:"postcode" json:"postcode"`
	InvoiceShippingPreference string `bson:"invoice_shipping_preference" json:"invoice_shipping_preference"`
	Address                   string `bson:"address" json:"address"`
	Status                    bool   `bson:"status" json:"status"`
	BillTo                    string `bson:"bill_to" json:"bill_to"`
}
