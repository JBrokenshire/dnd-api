package middleware

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"purplevisits.com/mdm/db/models"
	"purplevisits.com/mdm/api/responses"
)

// EapiMachine Will be called after enterprise_api-Key. Can pull machine ID and random out of headers to find the machine. Fail otherwise.
func EapiMachine(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			apiKey := c.Get("apiKey").(*models.ApiKey)
			machineId := c.Request().Header.Get("MachineId")
			random := c.Request().Header.Get("Random")

			var machine models.Machine
			db.Model(&machine).
				Where("enterprise_uid = ? AND machine_id = ? AND random = ?", apiKey.EnterpriseUid, machineId, random).
				First(&machine)

			if machine.ID == 0 {
				return responses.ErrorResponse(c, http.StatusUnauthorized, "Unable to find that machine")
			}

			c.Set("machine", machine)
			c.Set("random", random)

			return next(c)
		}
	}
}
