package models

type Product struct {
	ID       *int64  `json:"id,omitempty" bson:"id,omitempty"`
	Name     *string `json:"name,omitempty" bson:"name,omitempty"`
	Amount   *int64  `json:"amount,omitempty" bson:"amount,omitempty"`
	Count    *int64  `json:"count,omitempty" bson:"count,omitempty"`
	Image    *string `json:"image,omitempty" bson:"image,omitempty"`
	Currency *string `json:"currency,omitempty" bson:"currency,omitempty"`
}
