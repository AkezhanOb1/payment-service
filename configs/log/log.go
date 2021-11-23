package log

import (
	"time"

	"github.com/AkezhanOb1/payment/models"
	"github.com/sirupsen/logrus"
)

var Logger models.Logrus

func init() {
	var logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	Logger = models.Logrus{
		logger,
	}

}
