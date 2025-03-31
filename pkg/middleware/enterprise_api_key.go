package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"purplevisits.com/mdm/api/responses"
	"purplevisits.com/mdm/services"
)

func EnterpriseApiKey(ks *services.KeyService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Check endpoint is not being brute forced
			userIp := c.RealIP()
			key := c.Request().Header.Get("Authorization")
			daemonVersion := c.Request().Header.Get("DaemonVersion")
			apiKey, err := ks.EnterpriseKeyValidate(key, userIp)
			if err != nil {
				return responses.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("Key Issue: %v", err))
			}
			c.Set("apiKey", apiKey)
			c.Set("daemonVersion", daemonVersion)

			return next(c)
		}
	}
}
