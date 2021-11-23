package models

// Currency model
type Currency struct {
	Code       string `json:"code" bson:"code" validate:"required"`
	MinorUnits int    `json:"minorUnits" bson:"minorUnits" validate:"required"`
}
