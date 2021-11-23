package repositories

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func MerchantAPIRequestRepository(request []byte) (map[string]interface{}, error) {

	var apiMerchant = os.Getenv("ApiMerchant")

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("POST", apiMerchant, bytes.NewBuffer(request))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "ru")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	if _, ok := response["error"]; ok {
		e := response["error"].(map[string]interface{})
		return nil, fmt.Errorf("%v", e["message"])
	}
	return response["json"].(map[string]interface{}), nil
}
