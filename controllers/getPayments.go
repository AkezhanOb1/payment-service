package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/AkezhanOb1/payment/configs/log"
	"github.com/AkezhanOb1/payment/models"
	"github.com/AkezhanOb1/payment/services"
	"github.com/labstack/echo/v4"
)

//GetPaymentsController is a
func GetPaymentsController(c echo.Context) error {
	var start = time.Now()
	var customError models.CustomError
	var filters models.ParticipantsFilter

	var err = json.Unmarshal([]byte(c.QueryParam("filters")), &filters)
	if err != nil {
		customError = models.NewCustomError(
			http.StatusBadRequest,
			"can not find filters in query param",
		)
		log.Logger.Error(customError, c, http.StatusBadRequest, start)
		return c.JSON(http.StatusBadRequest, customError)
	}

	err = c.Validate(filters)
	if err != nil {
		customError = models.NewCustomError(
			http.StatusBadRequest,
			"can not validate provided fields"+err.Error(),
		)
		log.Logger.Error(customError, c, http.StatusBadRequest, start)
		return c.JSON(http.StatusBadRequest, customError)
	}

	if filters.Period.ToDate == nil {
		var t = time.Now().Format("2006-01-02T15:04:05")
		filters.Period.ToDate = &t
	}

	var sID = c.Get("sId").(string)
	participants, err := services.GetPaymentsService(sID, filters, context.Background())
	if err != nil {
		customError = models.NewCustomError(
			http.StatusInternalServerError,
			err.Error(),
		)
		log.Logger.Error(customError, c, http.StatusInternalServerError, start)
		log.Logger.Warn(customError)

		return c.JSON(http.StatusInternalServerError, customError)
	}

	log.Logger.OK(c, http.StatusOK, start)
	return c.JSON(200, participants)
}
