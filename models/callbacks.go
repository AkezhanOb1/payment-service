package models

//Callbacks is a
type Callbacks struct {
	ResultURL  *string `json:"resultUrl,omitempty" bson:"resultUrl,omitempty" validate:"omitempty"`
	SuccessURL string  `json:"successUrl" bson:"successUrl" validate:"required"`
	FailureURL string  `json:"failureUrl" bson:"failureUrl" validate:"required"`
	BackURL    *string `json:"backUrl,omitempty" bson:"backUrl,omitempty" validate:"omitempty"`
}
