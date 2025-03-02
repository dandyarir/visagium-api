package handler

import (
	"github.com/labstack/echo/v4"
)

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// RespondWithError responds with an error message
func RespondWithError(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ErrorResponse{
		Error: message,
	})
}
