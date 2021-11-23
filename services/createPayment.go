package services

import (
	"context"
	"os"
	"time"

	"github.com/AkezhanOb1/payment/models"
	"github.com/AkezhanOb1/payment/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePaymentService(payment models.Payment, ctx context.Context) (interface{}, error) {
	var status models.Status

	status.Code = "payment created"
	payment.Status = status
	payment.Id = primitive.NewObjectID()

	if payment.SID == "4449354148" {
		*payment.Options.Callbacks.ResultURL = "http://turan.team/guest.php/bloomz/result" + *payment.Options.Callbacks.ResultURL
	}

	siteName, err := repositories.GetMerchantSiteNameRepository(payment.SID, ctx)
	if err != nil {
		return nil, err
	}

	payment.SiteName = siteName
	payment.PaymentCreatedAt = time.Now().Format(time.RFC3339)

	_, err = repositories.CreatePaymentRepository(payment, ctx)
	if err != nil {
		return nil, err
	}

	var paymentPageURL = os.Getenv("PaymentPageURL")

	var paymentPage = models.PaymentPage{
		ID:             payment.Id,
		PaymentPageUrl: paymentPageURL + payment.Id.Hex(),
	}

	return paymentPage, nil
}
