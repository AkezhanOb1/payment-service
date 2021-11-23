package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/AkezhanOb1/payment/models"

	"github.com/AkezhanOb1/payment/repositories"
)

func ConfirmPaymentService(paymentID string, phoneNumber string, ctx context.Context) (interface{}, error) {

	payment, err := GetPaymentService(paymentID, ctx)
	if err != nil {
		return nil, err
	}

	if payment.Status.Code != "payment created" {
		return nil, fmt.Errorf("the payment was already confirmed")
	}

	session, err := merchantAuthService(payment.SID, ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Session", session)

	invoiceID, err := merchantCreateInvoice(*payment, session, phoneNumber, ctx)
	if err != nil {
		return nil, err
	}

	var newStatus models.Status
	newStatus.Code = "invoice created"

	payment.Status = newStatus
	payment.InvoiceID = invoiceID
	payment.PaymentRequest.PhoneNumber = phoneNumber
	payment.InvoiceCreatedAt = time.Now().Format(time.RFC3339)

	client, err := repositories.GetClientInfoRepository(phoneNumber, ctx)
	if err != nil {
		return nil, err
	}

	payment.Client = client

	err = repositories.GetPaymentByInvoiceIDRepository(invoiceID, ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = repositories.ChangePaymentStatusRepository(*payment, ctx)
			if err != nil {
				return nil, err
			}
			return payment, nil
		}
		return nil, err
	}

	return nil, fmt.Errorf("invoiceID - %s уже используется", invoiceID)
}
