package repositories

import (
	"context"

	"github.com/AkezhanOb1/payment/configs/mongo"
	"github.com/AkezhanOb1/payment/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckPaymentIdempotenceService is a
func CheckPaymentIdempotenceService(sID string, idempotence string, ctx context.Context) (*models.Payment, error) {
	var filter = bson.D{{"sId", sID}, {"idempotence", idempotence}}
	var result = mongo.Payments.FindOne(ctx, filter)
	var payment models.Payment

	err := result.Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
