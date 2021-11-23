package job

import (
	"context"
	"time"

	"github.com/AkezhanOb1/payment/configs/log"

	"github.com/AkezhanOb1/payment/services"

	"github.com/AkezhanOb1/payment/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func CancelPaymentsJob() {
	var now = time.Now()
	var past = now.Add(-10 * time.Minute).Format(time.RFC3339)
	var filter = bson.D{
		{"invoiceCreatedAt", bson.M{"$lte": past}},
		{"$or", []interface{}{
			bson.D{{"status.code", "invoice created"}},
			bson.D{{"status.code", "new"}},
		}},
		{"paymentRequest.forever", false},
	}
	payments, err := repositories.GetPaymentsWithFilter(context.Background(), filter)
	if err != nil {
		log.Logger.Warn(err)
	}

	log.Logger.Info("payment's to cancel from cron ", len(payments))
	for _, payment := range payments {

		status, err := services.MerchantGetInvoiceStatus(payment.Id.Hex(), context.Background())
		if err != nil {
			log.Logger.Warn(err)
			continue
		}

		log.Logger.Info("CRON STATUS", status)

		if status.Code == "approved" {
			continue
		}

		log.Logger.Info("canceling: ", payment.Id.Hex())
		_, err = services.CancelPaymentService(payment.Id.Hex(), context.Background())
		if err != nil {
			log.Logger.Warn(err)
			continue
		}

		log.Logger.Info("cron canceled payment id - ", payment.Id)
	}

}
