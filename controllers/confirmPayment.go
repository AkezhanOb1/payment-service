package controllers

import (
	"net/http"
	"time"

	"github.com/AkezhanOb1/payment/configs/log"
	"github.com/AkezhanOb1/payment/models"
	"github.com/AkezhanOb1/payment/services"
	"github.com/labstack/echo/v4"
)

func ConfirmPaymentController(c echo.Context) error {
	var start = time.Now()
	var customError models.CustomError

	var paymentID = c.Param("paymentID")
	if len(paymentID) != 24 {
		customError = models.NewCustomError(
			http.StatusBadRequest,
			"not valid paymentID provided",
		)
		log.Logger.Error(customError, c, http.StatusBadRequest, start)
		return c.JSON(http.StatusBadRequest, customError)
	}

	var phoneNumber = c.QueryParam("phoneNumber")
	if len(phoneNumber) != 11 {
		customError = models.NewCustomError(
			http.StatusBadRequest,
			"not valid phoneNumber in query param",
		)
		log.Logger.Error(customError, c, http.StatusBadRequest, start)
		return c.JSON(http.StatusBadRequest, customError)
	}

	payment, err := services.ConfirmPaymentService(paymentID, phoneNumber, c.Request().Context())
	if err != nil {
		customError = models.NewCustomError(
			http.StatusInternalServerError,
			err.Error(),
		)
		log.Logger.Error(customError, c, http.StatusInternalServerError, start)
		return c.JSON(http.StatusInternalServerError, customError)
	}

	log.Logger.OK(c, http.StatusOK, start)
	return c.JSON(http.StatusOK, payment)
}
