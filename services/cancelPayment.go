package services

import (
	"context"
	"fmt"

	"github.com/AkezhanOb1/payment/models"

	"github.com/AkezhanOb1/payment/repositories"
)

//CancelPaymentService is a
func CancelPaymentService(paymentID string, ctx context.Context) (interface{}, error) {

	payment, err := GetPaymentService(paymentID, ctx)
	if err != nil {
		return nil, err
	}

	if payment.Status.Code != "invoice created" && payment.Status.Code != "new" {
		return nil, fmt.Errorf("please confirm your payment first")
	}

	var status = models.Status{
		Code: "canceled",
	}

	payment.Status = status
	accountID, err := repositories.GetUserAccountId(payment.PhoneNumber, ctx)
	if err != nil {
		return nil, err
	}

	err = repositories.CancelInvoiceRepository(*payment, accountID, ctx)
	if err != nil {
		return nil, err
	}

	err = repositories.ChangePaymentStatusRepository(*payment, ctx)
	if err != nil {
		return nil, err
	}

	if payment.Options.Callbacks.ResultURL != nil {
		err = repositories.PutPaymentResultInQueue(*payment)
		if err != nil {
			return nil, err
		}
	}

	return payment, nil
}
