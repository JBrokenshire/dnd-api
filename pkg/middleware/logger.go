package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"os"
	"purplevisits.com/mdm/db/models"
	"strconv"
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

			//See if we have a current user and enterprise and device uids
			userUid := ""
			enterpriseUid := ""
			deviceUid := ""

			// FOr enterprise API calls, add in the id of the apikey and the enterprise.
			if strings.HasPrefix(v.URI, "/eapi/") {
				apiKeyI := c.Get("apiKey")
				if apiKeyI != nil {
					apiKey := apiKeyI.(*models.ApiKey)
					userUid = strconv.Itoa(int(apiKey.ID))
					enterpriseUid = apiKey.EnterpriseUid
				}
			} else if strings.HasPrefix(v.URI, "/machines/") {
				apiKeyI := c.Get("apiKey")
				if apiKeyI != nil {
					apiKey := apiKeyI.(*models.ApiKey)
					userUid = strconv.Itoa(int(apiKey.ID))
					enterpriseUid = apiKey.EnterpriseUid
				}
			} else {
				// See if we can find hte uid in the context and hte enterprise uid. These are set in Owner and User middlewares.
				uidI := c.Get("uid")
				if uidI != nil {
					userUid = uidI.(string)
				}

				enterpriseUidI := c.Get("enterpriseId")
				if enterpriseUidI != nil {
					enterpriseUid = enterpriseUidI.(string)
				}

				deviceUidI := c.Get("deviceUid")
				if deviceUidI != nil {
					deviceUid = deviceUidI.(string)
				}
			}

			logger.Info().
				Int("status", v.Status).
				Str("uri", v.URI).
				Str("method", v.Method).
				Str("user_agent", v.UserAgent).
				Str("user", userUid). // This may be UserUID, OwnerUID or Enterprise API KeyID.
				Str("enterprise", enterpriseUid).
				Str("device", deviceUid).
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
