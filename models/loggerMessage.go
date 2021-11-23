package models

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

//LoggerMessage is a
type LoggerMessage struct {
	Error   *CustomError `json:"omitempty"`
	SID     string
	Status  int
	Method  string
	URI     string
	Agent   string
	IP      string
	Latency string
}

func NewLoggerMessage(e *CustomError, status int, c echo.Context, latency string) LoggerMessage {
	return LoggerMessage{
		Error:   e,
		Status:  status,
		SID:     c.Get("sId").(string),
		Method:  c.Request().Method,
		URI:     c.Request().RequestURI,
		Agent:   c.Request().UserAgent(),
		IP:      c.Request().RemoteAddr,
		Latency: latency,
	}
}

func (lm LoggerMessage) String() string {
	if lm.Error != nil {
		return fmt.Sprintf(
			"error: %s, status: %d, method: %s, sid: %s, uri: %s, agent: %s, ip: %s, latency: %s",
			lm.Error.Error(),
			lm.Status,
			lm.Method,
			lm.SID,
			lm.URI,
			lm.Agent,
			lm.IP,
			lm.Latency,
		)
	}
	return fmt.Sprintf(
		"status: %d, method: %s, sid: %s, uri: %s, agent: %s, ip: %s, latency: %s",
		lm.Status,
		lm.Method,
		lm.SID,
		lm.URI,
		lm.Agent,
		lm.IP,
		lm.Latency,
	)
}
