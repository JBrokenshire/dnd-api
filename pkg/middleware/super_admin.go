package middleware

import (
	"github.com/labstack/echo/v4"
	"purplevisits.com/mdm/db/models"
)

// SuperAdmin Ensures the user is a super admin
func SuperAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("currentUser").(*models.User)

			if user.SuperAdmin == false {
				return errForbidden("Permission Error. Super Users Only")
			}

			return next(c)
		}
	}
}
