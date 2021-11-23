package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/AkezhanOb1/payment/models"
	"github.com/AkezhanOb1/payment/repositories"
)

func merchantCreateInvoice(payment models.Payment, session string, phoneNumber string, ctx context.Context) (string, error) {
	password, err := repositories.GetMerchantPasswordRepository(payment.SID, ctx)
	if err != nil {
		return "", err
	}

	var jsonMap = make(map[string]interface{})
	jsonMap = make(map[string]interface{})
	//jsonMap["lang"] = "ru"
	jsonMap["session"] = session
	jsonMap["mobile"] = phoneNumber
	jsonMap["amount"] = payment.Amount.Sum
	jsonMap["currency"] = payment.Amount.Currency.Code
	jsonMap["desc"] = payment.Description
	jsonMap["invoiceType"] = 1
	jsonMap["paymentTypeId"] = payment.PaymentType
	jsonMap["test"] = 0
	if payment.OrderID != nil {
		jsonMap["orderId"] = *payment.OrderID
	}

	if payment.Products != nil {
		var paymentCurrency = "KZT"
		for i := range payment.Products {
			if payment.Products[i].ID == nil {
				var id = int64(i + 1)
				payment.Products[i].ID = &id
			}

			if payment.Products[i].Count == nil {
				var count int64 = 1
				payment.Products[i].Count = &count
			}

			if payment.Products[i].Name == nil {
				payment.Products[i].Name = &payment.Description
			}
			payment.Products[i].Currency = &paymentCurrency
		}

		jsonMap["products"] = payment.Products
	}

	if payment.FieldsApps != nil {
		jsonMap["fieldsApp"] = payment.FieldsApps
	}

	if payment.CreateNotExistedUse != nil {
		jsonMap["createNotExistedUse"] = payment.CreateNotExistedUse
	}

	var apiAuthRequest = models.MerchantApiRequest{
		CMD:      "createInvoice",
		MKTime:   "1610551850",
		DateTime: time.Now().Format(time.RFC3339),
		SID:      payment.SID,
		JSON:     jsonMap,
	}

	jsonStr, _ := json.Marshal(apiAuthRequest)
	apiAuthRequest.Hash = hmacMD5(jsonStr, password)
	jsonStr, _ = json.Marshal(apiAuthRequest)

	response, err := repositories.MerchantAPIRequestRepository(jsonStr)
	if err != nil {
		return "", err
	}

	log.Println(response)
	invoiceID := response["invoiceId"].(string)

	return invoiceID, nil
}
