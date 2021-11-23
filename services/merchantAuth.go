package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/AkezhanOb1/payment/models"
	"github.com/AkezhanOb1/payment/repositories"
)

func merchantAuthService(sID string, ctx context.Context) (string, error) {

	//session, err := repositories.RetrieveMerchantApiSessionRepository(sID)
	//
	//if err != nil && err != redis.Nil {
	//	return "", err
	//}
	//
	//if session != "" {
	//	return session, err
	//}

	password, err := repositories.GetMerchantPasswordRepository(sID, ctx)
	if err != nil {
		return "", err
	}

	pin, err := repositories.GetMerchantPinRepository(sID, ctx)
	if err != nil {
		return "", err
	}

	var jsonMap = make(map[string]interface{})
	jsonMap["pin"] = hash256SHA(hashMD5(pin))
	var apiAuthRequest = models.MerchantApiRequest{
		CMD:      "authAgent",
		MKTime:   "1610551850",
		DateTime: time.Now().Format(time.RFC3339),
		SID:      sID,
		JSON:     jsonMap,
	}
	jsonStr, _ := json.Marshal(apiAuthRequest)
	apiAuthRequest.Hash = hmacMD5(jsonStr, password)
	jsonStr, _ = json.Marshal(apiAuthRequest)

	response, err := repositories.MerchantAPIRequestRepository(jsonStr)
	if err != nil {
		return "", err
	}

	session := response["session"].(string)
	//err = repositories.SetMerchantApiSessionRepository(sID, session)
	//if err != nil {
	//	return "", err
	//}
	return session, nil
}
