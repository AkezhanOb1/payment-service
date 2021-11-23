package models

import (
	"fmt"
)

//CustomError is a
type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", err.Code, err.Message)
}

func NewCustomError(code int, message string) CustomError {
	return CustomError{
		Code:    code,
		Message: message,
	}
}

//func (err CustomError) String(c echo.Context, start time.Time) string {
//	return NewLoggerMessage(&err, http.StatusBadRequest, c, time.Since(start).String()).String()
//}
