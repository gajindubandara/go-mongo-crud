package models

type Pet struct {
	ID           string  `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string  `bson:"name" json:"name"`
	Color        string  `bson:"color" json:"color"`
	Breed        string  `bson:"breed" json:"breed"`
	DateOfBirth  string  `bson:"date_of_birth" json:"date_of_birth"`
	SellingPrice float32 `bson:"selling_price" json:"selling_price"`
}
