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

type Room struct {
	ID            string `bson:"_id,omitempty" json:"id"`
	BuildingID    string `bson:"building_id" json:"building_id"`
	FloorNumber   string `bson:"floor_number" json:"floor_number"`
	FloorSquare   string `bson:"floor_square" json:"floor_square"`
	Buitenzijde   string `bson:"buitenzijde" json:"buitenzijde"`
	RoomType      string `bson:"room_type" json:"room_type"`
	WallType      string `bson:"wall_type" json:"wall_type"`
	RoomNumber    string `bson:"room_number" json:"room_number"`
	Binnenzijde   string `bson:"binnenzijde" json:"binnenzijde"`
	SeperstieGlas string `bson:"seperstie_glas" json:"seperstie_glas"`
	UsageType     string `bson:"usage_type" json:"usage_type"`
	FloorType     string `bson:"floor_type" json:"floor_type"`
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
