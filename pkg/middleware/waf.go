package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"purplevisits.com/mdm/pkg/waf"
	"purplevisits.com/mdm/api/responses"
)

// WafMiddleware is a middleware that provides a Web Application Firewall (WAF) for the application.
// It uses the Coraza WAF library to provide the WAF functionality.

//Add echo middleware using Coraza WAF

func WafMiddleware() echo.MiddlewareFunc {
	wafService := waf.NewService()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := wafService.ProcessRequest(c.Request())
			if err != nil {
				return responses.ErrorResponse(c, http.StatusUnauthorized, "Error processing request (7413)")
			}
			return next(c)
		}
	}
}
