package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/AkezhanOb1/payment/services"

	"github.com/AkezhanOb1/payment/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

//TestGetPaymentService is a
func TestGetPaymentService(t *testing.T) {
	var filter = bson.D{}

	payments, err := repositories.GetPaymentsWithFilter(context.Background(), filter)
	if err != nil {
		t.Error(err)
	}

	for _, payment := range payments {
		paymentInfo, err := services.GetPaymentService(payment.Id.Hex(), context.Background())
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(payment, paymentInfo) == true {
			t.Error("different struct")
		}
	}
}
