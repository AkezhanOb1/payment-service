package services

import (
	"encoding/json"
	"log"
	"time"

	"context"

	"github.com/AkezhanOb1/payment/models"
	"github.com/AkezhanOb1/payment/repositories"
)

//TODO: put password for api in context
func MerchantGetInvoiceStatus(paymentID string, ctx context.Context) (*models.Status, error) {

	payment, err := GetPaymentService(paymentID, ctx)
	if err != nil {
		return nil, err
	}

	password, err := repositories.GetMerchantPasswordRepository(payment.SID, ctx)
	if err != nil {
		return nil, err
	}

	session, err := merchantAuthService(payment.SID, ctx)
	if err != nil {
		return nil, err
	}

	var jsonMap = make(map[string]interface{})
	jsonMap = make(map[string]interface{})
	jsonMap["session"] = session
	jsonMap["invoiceId"] = payment.InvoiceID

	var apiAuthRequest = models.MerchantApiRequest{
		CMD:      "getStatus",
		MKTime:   "1612156723741",
		DateTime: time.Now().Format(time.RFC3339),
		SID:      payment.SID,
		JSON:     jsonMap,
	}

	jsonStr, _ := json.Marshal(apiAuthRequest)
	apiAuthRequest.Hash = hmacMD5(jsonStr, password)
	jsonStr, _ = json.Marshal(apiAuthRequest)

	response, err := repositories.MerchantAPIRequestRepository(jsonStr)
	if err != nil {
		return nil, err
	}
	var newStatus models.Status
	var invoice = response["invoiceStatus"].(map[string]interface{})
	if _, ok := invoice["payments"]; ok {
		var payments = invoice["payments"].(interface{})
		newStatus.Code = payments.(map[string]interface{})["status"].(string)
		if payment.Status != newStatus {
			payment.Status = newStatus
			err = repositories.ChangePaymentStatusRepository(*payment, ctx)
			if err != nil {
				return nil, err
			}
			if newStatus.Code == "approved" || newStatus.Code == "canceled" {
				if payment.Options.Callbacks.ResultURL != nil {
					if newStatus.Code == "approved" {
						operationID, err := repositories.GetOperationIdRepository(payment.InvoiceID, ctx)
						if err != nil {
							log.Println("error in operationIdRep:", err)
						}
						payment.OperationID = operationID
						err = repositories.ChangePaymentStatusRepository(*payment, ctx)
						if err != nil {
							log.Println("error in changing operationID status :", err)
						}
					}
					err = repositories.PutPaymentResultInQueue(*payment)
					if err != nil {
						return nil, err
					}
				}
			}
		}
	} else {
		var status = invoice["status"].(float64)
		if status == 0 {
			newStatus.Code = "invoice created"
		}
	}

	return &newStatus, nil
}
