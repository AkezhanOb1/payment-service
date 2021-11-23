package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentPage struct {
	ID             primitive.ObjectID `json:"id"`
	PaymentPageUrl string             `json:"paymentPageUrl"`
}
