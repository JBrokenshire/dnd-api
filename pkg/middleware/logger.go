package middleware

import (
	"dnd-api/db/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"os"
)

var CustomLogFunction func(v middleware.RequestLoggerValues)

func CustomLogger() echo.MiddlewareFunc {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper: func(c echo.Context) bool { return false },

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

			//See if we have a current user and enterprise and device uids
			var userId uint

			if c.Get("currentUser") != nil {
				currentUser := c.Get("currentUser").(*models.User)
				userId = currentUser.ID
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
