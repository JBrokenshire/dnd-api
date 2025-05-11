package middleware

import (
	m "dnd-api/db/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"os"
	"strings"
)

var CustomLogFunction func(v middleware.RequestLoggerValues)

func CustomLogger() echo.MiddlewareFunc {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper: func(c echo.Context) bool {
			// Functions to skip the login

			// If the user agent is ELB-HealthChecker then skip
			if strings.HasPrefix(c.Request().UserAgent(), "ELB-HealthChecker") {
				return true
			}
			return false
		},
		LogStatus:        true,
		LogURI:           true,
		LogMethod:        true,
		LogUserAgent:     true,
		LogHost:          true,
		LogError:         true,
		LogLatency:       true,
		LogRemoteIP:      true,
		LogResponseSize:  true,
		LogContentLength: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var userId uint
			currentUser := c.Get("currentUser")
			if currentUser != nil {
				currentUser = currentUser.(*m.User)
				userId = currentUser.(*m.User).ID
			}

			logger.Info().
				Int("status", v.Status).
				Str("uri", v.URI).
				Str("method", v.Method).
				Str("user_agent", v.UserAgent).
				Uint("user", userId).
				Str("host", v.Host).
				Dur("latency", v.Latency).
				Str("latency_human", v.Latency.String()).
				Str("remote_ip", v.RemoteIP).
				Str("bytes_in", v.ContentLength).
				Int64("bytes_out", v.ResponseSize).
				Err(v.Error).
				Msg("")

			if CustomLogFunction != nil {
				CustomLogFunction(v)
			}

			return nil
		},
	})
}
