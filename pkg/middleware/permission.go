package middleware

import (
	"github.com/labstack/echo/v4"
	models2 "purplevisits.com/mdm/db/models"
)

type Permission struct {
	Subject models2.Subject
	Action  models2.Action
}
type Permissions []Permission

// HasPermission will check that the current user has the permission specified.
func HasPermission(s models2.Subject, a models2.Action) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			currentUser := c.Get("currentUser").(*models2.User)
			if currentUser == nil {
				return errForbidden("NCU: Invalid Permissions")
			}
			if currentUser.HasPermission(s, a) == false {
				return errForbidden("Missing Required Permission")
			}
			return next(c)
		}
	}
}

func HasOneOfPermissions(permissions Permissions) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			currentUser := c.Get("currentUser").(*models2.User)
			if currentUser == nil {
				return errForbidden("NCU: Invalid Permissions")
			}
			hasOneOf := false
			for _, permission := range permissions {
				if currentUser.HasPermission(permission.Subject, permission.Action) == true {
					hasOneOf = true
				}
			}

			if hasOneOf == false {
				return errForbidden("Missing Required Permission")
			}
			return next(c)
		}
	}
}
