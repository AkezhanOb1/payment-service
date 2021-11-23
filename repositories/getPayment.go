package repositories

import (
	"context"

	"github.com/AkezhanOb1/payment/configs/mongo"
	"github.com/AkezhanOb1/payment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPaymentRepository(paymentID string, ctx context.Context) (*models.Payment, error) {
	var objectID, _ = primitive.ObjectIDFromHex(paymentID)
	var filter = bson.D{{"_id", objectID}}
	var result = mongo.Payments.FindOne(ctx, filter)

	var payment models.Payment

	err := result.Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
