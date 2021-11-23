package models

type Client struct {
	User User       `json:"user"`
	KYC  *Reference `json:"kyc,omitempty"`
}
