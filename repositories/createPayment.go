package repositories

import (
	"context"

	"github.com/AkezhanOb1/payment/configs/mongo"
	"github.com/AkezhanOb1/payment/models"
)

func CreatePaymentRepository(payment models.Payment, ctx context.Context) (interface{}, error) {
	_, err := mongo.Payments.InsertOne(ctx, payment)
	if err != nil {
		return payment, err
	}

	return payment, nil
}
