package middleware

import "github.com/labstack/echo/v4"

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		return next(c)
	}
}
