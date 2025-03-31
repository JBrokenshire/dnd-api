package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func ValidateCronCode(next echo.HandlerFunc) echo.HandlerFunc {
	validCode := os.Getenv("CRON_CODE_KEY")
	return func(c echo.Context) error {

		if validCode == "" {
			log.Println("Cron code not set in ENV!")
			return errForbidden("Server not set up for Cron yet")
		}

		cronCode := c.Request().Header.Get("CronCode")

		if validCode != cronCode {
			log.Printf("Invalid cron code passed through. Code Length: %v", len(cronCode))
			return errForbidden("Invalid Code in request")
		}
		return next(c)
	}
}
