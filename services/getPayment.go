package services

import (
	"context"

	"github.com/AkezhanOb1/payment/models"
	"github.com/AkezhanOb1/payment/repositories"
)

func GetPaymentService(paymentID string, ctx context.Context) (*models.Payment, error) {
	payment, err := repositories.GetPaymentRepository(paymentID, ctx)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
