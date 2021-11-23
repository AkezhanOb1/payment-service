package models

//ParticipantsFilter is a
type ParticipantsFilter struct {
	Period  Period   `json:"period" validate:"required"`
	Amounts *Amounts `json:"amounts,omitempty"`
	Status  *string  `json:"status,omitempty"`
}
