package repositories

import (
	"context"

	"github.com/AkezhanOb1/payment/configs/mongo"
	"github.com/AkezhanOb1/payment/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPaymentByInvoiceIDRepository(invoiceID string, ctx context.Context) error {
	var filter = bson.D{{"invoiceId", invoiceID}}
	var result = mongo.Payments.FindOne(ctx, filter)

	var payment models.Payment

	err := result.Decode(&payment)
	if err != nil {
		return err
	}

	return nil
}
