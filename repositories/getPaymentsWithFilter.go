package repositories

import (
	"context"

	"github.com/AkezhanOb1/payment/configs/mongo"

	"github.com/AkezhanOb1/payment/models"
)

func GetPaymentsWithFilter(ctx context.Context, filter interface{}) ([]models.Payment, error) {

	cursor, err := mongo.Payments.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var payments []models.Payment

	if err = cursor.All(ctx, &payments); err != nil {
		return nil, err
	}

	return payments, nil
}
