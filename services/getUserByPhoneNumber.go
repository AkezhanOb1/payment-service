package services

import (
	"context"

	"github.com/AkezhanOb1/payment/repositories"
)

func GetUserByPhoneNumberService(phoneNumber string, ctx context.Context) (interface{}, error) {

	user, err := repositories.GetUserByPhoneNumberRepository(phoneNumber, ctx)
	if err != nil {
		return nil, err
	}
	return user, nil

}
