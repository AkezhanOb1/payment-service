package services

import (
	"context"

	"github.com/AkezhanOb1/payment/models"
	"github.com/AkezhanOb1/payment/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

//CheckPaymentIdempotenceService is a
func CheckPaymentIdempotenceService(sID string, idempotence string, ctx context.Context) (*models.Payment, error) {
	resp, err := repositories.CheckPaymentIdempotenceService(sID, idempotence, ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return resp, err
}
