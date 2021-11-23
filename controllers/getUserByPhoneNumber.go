package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/AkezhanOb1/payment/configs/log"

	"github.com/AkezhanOb1/payment/services"

	"github.com/AkezhanOb1/payment/models"
	"github.com/labstack/echo/v4"
)

//GetUserByPhoneNumberController is a
func GetUserByPhoneNumberController(c echo.Context) error {
	var start = time.Now()
	var customError models.CustomError
	var phoneNumber = c.QueryParam("phoneNumber")
	if phoneNumber == "" || len(phoneNumber) != 11 {
		customError = models.NewCustomError(
			http.StatusBadRequest,
			"not valid phone number provided",
		)

		log.Logger.Error(customError, c, http.StatusBadRequest, start)
		return c.JSON(http.StatusBadRequest, customError)
	}

	user, err := services.GetUserByPhoneNumberService(phoneNumber, context.Background())
	if err != nil {
		customError = models.NewCustomError(
			http.StatusInternalServerError,
			err.Error(),
		)
		log.Logger.Error(customError, c, http.StatusInternalServerError, start)
		return c.JSON(http.StatusInternalServerError, customError)
	}

	log.Logger.OK(c, http.StatusOK, start)
	return c.JSON(http.StatusOK, user)
}
