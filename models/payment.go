package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Payment is a
type Payment struct {
	Id               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SID              string             `json:"sId,omitempty" bson:"sId,omitempty"`
	SiteName         string             `json:"siteName,omitempty" bson:"siteName,omitempty"`
	Idempotence      string             `json:"idempotence,omitempty" bson:"idempotence,omitempty"`
	InvoiceID        string             `json:"invoiceId,omitempty" bson:"invoiceId,omitempty"`
	InvoiceCreatedAt string             `json:"invoiceCreatedAt,omitempty" bson:"invoiceCreatedAt,omitempty"`
	PaymentCreatedAt string             `json:"paymentCreatedAt,omitempty" bson:"paymentCreatedAt,omitempty"`
	Status           Status             `json:"status,omitempty" bson:"status,omitempty"`
	Client           *User              `json:"client,omitempty" bson:"client,omitempty"`
	OperationID      string             `json:"operationId,omitempty" bson:"operationId,omitempty"`
	PaymentRequest   `bson:"paymentRequest"`
}

type PaymentRequest struct {
	PaymentType         int         `json:"paymentType" bson:"paymentType" validate:"required"`
	Amount              Amount      `json:"amount" bson:"amount" validate:"dive,required"`
	Description         string      `json:"description" bson:"description" validate:"required"`
	OrderID             *string     `json:"orderId,omitempty" bson:"orderId,omitempty"`
	PhoneNumber         string      `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty" validate:"omitempty,len=11"`
	Options             Options     `json:"options" bson:"options" validate:"dive,required"`
	Language            string      `json:"language,omitempty"  bson:"language,omitempty"`
	Products            []Product   `json:"products,omitempty" bson:"products,omitempty"`
	FieldsApps          []FieldsApp `json:"fieldsApp,omitempty" bson:"fieldsApp,omitempty"`
	CreateNotExistedUse *bool       `json:"createNotExistedUse,omitempty" bson:"createNotExistedUse,omitempty"`
	Forever             bool        `json:"forever" bson:"forever"`
}
