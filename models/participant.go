package models

//Participant is a
type Participant struct {
	TransactionID int64   `json:"transactionID"`
	UserID        int64   `json:"userID"`
	WalletID      int64   `json:"walletID"`
	PhoneNumber   string  `json:"phoneNumber"`
	Amount        float64 `json:"amount"`
	CreatedAt     string  `json:"createdAt"`
	IIN           *string `json:"iin,omitempty"`
	FirstName     *string `json:"firstName,omitempty"`
	SecondName    *string `json:"secondName,omitempty"`
	Patronymic    *string `json:"patronymic,omitempty"`
}

type ParticipantsList struct {
	Participants []Participant `json:"participants"`
}
