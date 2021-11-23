package tests

import (
	"context"
	"log"
	"testing"

	"github.com/AkezhanOb1/payment/models"

	"github.com/AkezhanOb1/payment/services"

	"github.com/AkezhanOb1/payment/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

//TestCancelPaymentService is a
func TestCancelPaymentService(t *testing.T) {
	var filter = bson.D{}

	payments, err := repositories.GetPaymentsWithFilter(context.Background(), filter)
	if err != nil {
		t.Error(err)
	}

	log.Println(len(payments))

	for i, payment := range payments {
		res, err := services.CancelPaymentService(payment.Id.Hex(), context.Background())
		if err != nil {
			switch err.Error() {
			case " please confirm your payment first":
				t.Log(i, ":please confirm your payment first", payment.Id.Hex(), payment.Status.Code)
			default:
				t.Error(i, ":", err)
			}
			continue
		}

		p, ok := res.(*models.Payment)
		if !ok {
			t.Error(i, ": can not cast to payment struct")
		}

		if p.Status.Code != "canceled" {
			t.Error(i, ": status is another", p.Status.Code)
		}
		t.Log(i, ": OK")
	}
}
