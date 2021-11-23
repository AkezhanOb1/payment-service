package models

// Amount struct
type Amount struct {
	Sum           int64    `json:"sum"  bson:"sum" validate:"required"`
	Currency      Currency `json:"currency" bson:"currency" validate:"dive,required"`
	SumBeautified *string  `json:"sumBeautified,omitempty" bson:"sumBeautified,omitempty"`
}

type Amounts struct {
	BottomAmount *Amount `json:"bottomAmount,omitempty"`
	TopAmount    *Amount `json:"topAmount,omitempty"`
}
