package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Respond with a forbidden message
func errForbidden(message string) error {
	return &echo.HTTPError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

// Respond with a Not found message
func errNotFound(message string) error {
	return &echo.HTTPError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func errUnauthorized(message string) error {
	return &echo.HTTPError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}
