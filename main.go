package main

import (
	"net/http"

	_ "github.com/AkezhanOb1/payment/configs"
	"github.com/AkezhanOb1/payment/configs/log"
	"github.com/AkezhanOb1/payment/constants"
	"github.com/AkezhanOb1/payment/controllers"
	"github.com/AkezhanOb1/payment/job"
	"github.com/AkezhanOb1/payment/middlewares"
	"github.com/AkezhanOb1/payment/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
)

func main() {
	var e = echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodPost, http.MethodGet},
	}))
	e.Use(middleware.BasicAuth(middlewares.AuthMiddleware))
	e.Validator = &models.Validator{Validator: validator.New()}

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	c := cron.New()
	c.AddFunc("*/10 * * * *", job.CancelPaymentsJob)
	go c.Start()
	go job.ConsumeFromPaymentResultJob()

	e.GET("/user", controllers.GetUserByPhoneNumberController)
	e.POST("/payments", controllers.CreatePaymentsController)
	e.GET("/payments", controllers.GetPaymentsController)
	e.GET("/payments/:paymentID", controllers.GetPaymentController)
	e.POST("/payments/:paymentID/invoice", controllers.ConfirmPaymentController)
	e.POST("/payments/:paymentID/cancel", controllers.CancelPaymentController)

	log.Logger.Fatal(e.Start(constants.Port))
}
