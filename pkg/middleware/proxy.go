package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func ValidateProxyCode(next echo.HandlerFunc) echo.HandlerFunc {
	validCode := os.Getenv("PROXY_CODE_KEY")
	return func(c echo.Context) error {

		if validCode == "" {
			log.Println("Proxy code not set in ENV!")
			return errForbidden("Server not set up for Proxy yet")
		}

		proxyCode := c.Request().Header.Get("ProxyCode")

		if validCode != proxyCode {
			log.Printf("Invalid proxy code passed through. Code Length: %v", len(proxyCode))
			return errForbidden("Invalid Code in request")
		}
		return next(c)
	}
}
