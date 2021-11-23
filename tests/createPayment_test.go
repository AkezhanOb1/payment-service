package tests

import (
	"context"
	"testing"

	"github.com/AkezhanOb1/payment/services"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/AkezhanOb1/payment/models"
)

//TestPaymentService is a
func TestCreatePaymentService(t *testing.T) {
	gofakeit.Seed(6)

	var payments []models.Payment
	for i := 0; i < 150; i++ {
		var payment = models.Payment{
			SID:         "1111111111",
			Idempotence: gofakeit.UUID(),
			PaymentType: 1,
			PhoneNumber: gofakeit.Phone(),
			Amount: models.Amount{
				Sum: int64(gofakeit.Price(10.0, 1500000.5555)),
				Currency: models.Currency{
					Code:       "KZT",
					MinorUnits: 100,
				},
			},
			Description: gofakeit.HipsterSentence(3),
			Options: models.Options{
				Callbacks: models.Callbacks{
					ResultURL:  gofakeit.URL(),
					SuccessURL: gofakeit.URL(),
					FailureURL: gofakeit.URL(),
				},
			},
			Language: gofakeit.LanguageAbbreviation(),
		}

		payments = append(payments, payment)
	}

	for _, payment := range payments {
		res, err := services.CreatePaymentService(payment, context.Background())
		if err != nil {
			t.Error(err)
		}

		_, ok := res.(models.PaymentPage)
		if !ok {
			t.Error("can not cast to payment page struct")
		}
	}
}
