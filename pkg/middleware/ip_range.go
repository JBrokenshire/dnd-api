package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"purplevisits.com/mdm/pkg/ip"
	"purplevisits.com/mdm/api/responses"
)

func IpRange(cidrRanges string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			usersIp := c.RealIP()
			if inRange := ip.InRanges(usersIp, cidrRanges); inRange == false {
				return responses.ErrorResponse(c, http.StatusBadRequest, "Invalid IP address for KEY used")
			}

			return next(c)
		}
	}
}
