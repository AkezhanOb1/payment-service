package repositories

import (
	"context"

	"github.com/AkezhanOb1/payment/configs/mongo"
	"github.com/AkezhanOb1/payment/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChangePaymentStatusRepository is a
func ChangePaymentStatusRepository(payment models.Payment, ctx context.Context) error {
	_, err := mongo.Payments.ReplaceOne(
		ctx,
		bson.M{"_id": payment.Id},
		payment,
	)
	if err != nil {
		return err
	}
	return nil
}
