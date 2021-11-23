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

//TestConfirmPaymentService is a
func TestConfirmPaymentService(t *testing.T) {
	var filter = bson.D{}

	payments, err := repositories.GetPaymentsWithFilter(context.Background(), filter)
	if err != nil {
		t.Error(err)
	}

	log.Println(len(payments))

	var phoneNumbers = []string{"77772059339", "77781545994", "12312312312"}
	for i, payment := range payments {
		res, err := services.ConfirmPaymentService(payment.Id.Hex(), phoneNumbers[i%len(phoneNumbers)], context.Background())
		if err != nil {
			switch err.Error() {
			case "the payment was already confirmed":
				t.Log(i, ": the payment was already confirmed", payment.Id.Hex(), payment.Status.Code)
			case "User not found":
				t.Log(i, ": User not found:", payment.Id.Hex())
			default:
				t.Error(err)
			}
			continue
		}

		_, ok := res.(*models.Payment)
		if !ok {
			t.Error("can not cast to payment struct")
		}
		t.Log(i, ": OK")
	}
}
