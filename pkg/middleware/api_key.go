package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func APIKey(ks *services.KeyService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Check endpoint is not being brute forced
			userIp := c.RealIP()
			key := c.Request().Header.Get("Authorization")
			if err := ks.ExternalKeyOk(key, userIp); err != nil {
				return responses.ErrorResponse(c, http.StatusUnauthorized, err.Error())
			}
			return next(c)
		}
	}
}
