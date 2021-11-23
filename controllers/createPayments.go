package controllers

import (
	"net/http"
	"time"

	"github.com/AkezhanOb1/payment/configs/log"
	"github.com/AkezhanOb1/payment/models"
	"github.com/AkezhanOb1/payment/services"
	"github.com/labstack/echo/v4"
)

func CreatePaymentsController(c echo.Context) error {
	var start = time.Now()
	var customError models.CustomError
	var paymentRequest models.PaymentRequest

	var err = c.Bind(&paymentRequest)
	if err != nil {
		customError = models.NewCustomError(
			http.StatusBadRequest,
			err.Error(),
		)
		log.Logger.Error(customError, c, http.StatusBadRequest, start)
		return c.JSON(http.StatusBadRequest, customError)
	}

	err = c.Validate(paymentRequest)
	if err != nil {
		customError = models.NewCustomError(
			http.StatusBadRequest,
			"can not validate provided fields"+err.Error(),
		)
		log.Logger.Error(customError, c, http.StatusBadRequest, start)
		return c.JSON(http.StatusBadRequest, customError)
	}

	var idempotence = c.Request().Header.Get("X-Idempotency")
	var ok = validateUUID(idempotence)
	if ok == false {
		customError = models.NewCustomError(
			http.StatusBadRequest,
			"not valid uuid in X-Idempotency header",
		)
		log.Logger.Error(customError, c, http.StatusBadRequest, start)
		return c.JSON(http.StatusBadRequest, customError)
	}

	var payment models.Payment
	payment.PaymentRequest = paymentRequest
	payment.Idempotence = idempotence
	payment.SID = c.Get("sId").(string)

	prevPayment, err := services.CheckPaymentIdempotenceService(payment.SID, idempotence, c.Request().Context())
	if err != nil {
		customError = models.NewCustomError(
			http.StatusInternalServerError,
			err.Error(),
		)
		log.Logger.Error(customError, c, http.StatusInternalServerError, start)
		return c.JSON(http.StatusInternalServerError, customError)
	}

	if prevPayment != nil {
		log.Logger.OK(c, http.StatusOK, start)
		return c.JSON(http.StatusOK, prevPayment)
	}

	paymentPage, err := services.CreatePaymentService(payment, c.Request().Context())
	if err != nil {
		customError = models.NewCustomError(
			http.StatusInternalServerError,
			err.Error(),
		)
		log.Logger.Error(customError, c, http.StatusInternalServerError, start)
		return c.JSON(http.StatusInternalServerError, customError)
	}

	log.Logger.OK(c, http.StatusOK, start)
	return c.JSON(http.StatusOK, paymentPage)
}
