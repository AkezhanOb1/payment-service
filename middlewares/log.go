package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AkezhanOb1/payment/configs/log"
	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
)

func makeLogEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return log.Logger.WithFields(logrus.Fields{
			"at": time.Now().Format(time.RFC3339),
		})
	}

	return log.Logger.WithFields(logrus.Fields{
		"at":     time.Now().Format(time.RFC3339),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Logger.Info(fmt.Printf(""))
		makeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}
