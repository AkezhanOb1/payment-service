package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AkezhanOb1/payment/models"
)

func SendResultToMerchantRepository(payment models.Payment) error {
	var url = *payment.Options.Callbacks.ResultURL

	var jsonStr, _ = json.Marshal(payment)

	log.Println("URL:", url, "BODY:", string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("the request to the merchant was not succesfull, the status code is: %d", resp.StatusCode)
	}

	log.Println("RESP STATUS", resp.StatusCode)

	return nil
}
