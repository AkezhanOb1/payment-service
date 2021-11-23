package models

type User struct {
	ID          *int64  `json:"id,omitempty" bson:"id,omitempty"`
	FirstName   *string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName    *string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Patronymic  *string `json:"patronymic,omitempty" bson:"patronymic,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	IIN         *string `json:"iin,omitempty" bson:"iin,omitempty"`
	WalletID    *int64  `json:"walletId,omitempty" bson:"walletId,omitempty"`
}
