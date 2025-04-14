package middleware

import (
	"dnd-api/api/responses"
	"dnd-api/pkg/security"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Security checks the useragent and query params for anything dodgy
func Security(allowBlank bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userAgent := c.Request().UserAgent()

			if security.IsUserAgentBad(userAgent, allowBlank) {
				return responses.ErrorResponse(c, http.StatusUnauthorized, "Error processing request (8492)")
			}

			fullUrl := c.Request().URL.String()
			if security.IsURLBad(fullUrl) {
				return responses.ErrorResponse(c, http.StatusUnauthorized, "Error processing request (2489)")
			}

			return next(c)
		}
	}
}
