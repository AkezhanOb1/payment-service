package models

//Sum is a
type Sum struct {
	FromSum Amount  `json:"fromSum" validate:"omitempty,gte=0"`
	ToSum   *Amount `json:"toSum" validate:"omitempty,gte=0"`
}
