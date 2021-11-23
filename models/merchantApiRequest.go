package models

//MerchantApiRequest is a
type MerchantApiRequest struct {
	CMD      string                 `json:"cmd"`
	MKTime   string                 `json:"mktime"`
	DateTime string                 `json:"datetime"`
	SID      string                 `json:"sid"`
	JSON     map[string]interface{} `json:"json"`
	Hash     string                 `json:"hash,omitempty"`
}
