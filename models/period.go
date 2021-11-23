package models

//Period is a
type Period struct {
	FromDate string  `json:"fromDate" validate:"required"`
	ToDate   *string `json:"toDate,omitempty" validate:"omitempty"`
}
