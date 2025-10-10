package models

type Product struct {
	ID    string  `bson:"_id,omitempty" json:"id"`
	Name  string  `bson:"name" json:"name"`
	Price float64 `bson:"price" json:"price"`
}

type ProductWriteBody struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
