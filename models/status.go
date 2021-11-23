package models

//Status is
type Status struct {
	Code  string `json:"code"`
	Error string `json:"error_description,omitempty" bson:"error_description,omitempty"`
}
