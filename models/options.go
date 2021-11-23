package models

//Options is a
type Options struct {
	Callbacks   Callbacks `json:"callbacks" bson:"callbacks" validate:"dive,required"`
	PublicOffer *string   `json:"publicOffer,omitempty" bson:"publicOffer,omitempty" validate:"omitempty,url"`
	PromoCode   *string   `json:"promoCode,omitempty" bson:"promoCode,omitempty" validate:"omitempty"`
}
