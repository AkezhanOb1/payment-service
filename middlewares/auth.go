package middlewares

import (
	"net/http"

	"github.com/AkezhanOb1/payment/configs/log"

	"github.com/AkezhanOb1/payment/repositories"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(sID string, password string, c echo.Context) (bool, error) {
	req := c.Request()
	apiPassword, err := repositories.GetMerchantPasswordRepository(sID, c.Request().Context())
	if err != nil {
		return false, nil
	}

	if apiPassword != password {
		log.Logger.WithFields(map[string]interface{}{
			"method": req.Method,
			"uri":    req.RequestURI,
			"status": http.StatusUnauthorized,
			"ip":     req.RemoteAddr,
			"agent":  req.UserAgent(),
			"sid":    sID,
		}).Info()
		return false, nil
	}

	c.Set("sId", sID)
	return true, nil
}
