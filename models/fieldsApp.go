package models

type FieldsApp struct {
	FieldsName      *string `json:"fieldsName,omitempty" bson:"fieldsName,omitempty"`
	FieldsType      *string `json:"fieldsType,omitempty" bson:"fieldsType,omitempty"`
	FieldsMaxLenght *int64  `json:"fieldsMaxLenght,omitempty" bson:"fieldsMaxLenght,omitempty"`
	FieldsMinLenght *int64  `json:"fieldsMinLenght,omitempty" bson:"fieldsMinLenght,omitempty"`
}
