package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/AkezhanOb1/payment/repositories"

	"github.com/AkezhanOb1/payment/models"
)

func GetPaymentsService(siteID string, filters models.ParticipantsFilter, ctx context.Context) ([]models.Payment, error) {

	var filter = []bson.M{
		{"sId": siteID},
		{"invoiceCreatedAt": bson.M{"$gte": filters.Period.FromDate}},
		{"invoiceCreatedAt": bson.M{"$lte": filters.Period.ToDate}},
	}

	if filters.Amounts != nil {
		if filters.Amounts.BottomAmount != nil {
			filter = append(filter, bson.M{"paymentRequest.amount.sum": bson.M{"$gte": filters.Amounts.BottomAmount.Sum}})
		}
		if filters.Amounts.TopAmount != nil {
			filter = append(filter, bson.M{"paymentRequest.amount.sum": bson.M{"$lte": filters.Amounts.TopAmount.Sum}})
		}
	}

	if filters.Status != nil {
		filter = append(filter, bson.M{"status.code": filters.Status})
	}

	var filterQuery = bson.M{
		"$and": filter,
	}

	payments, err := repositories.GetPaymentsWithFilter(context.Background(), filterQuery)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
