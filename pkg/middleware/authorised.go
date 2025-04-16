package middleware

import (
	"dnd-api/services/jwt_service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Authorised check that the JWT flag of Authorised is True
func Authorised(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwt_service.JwtCustomClaims)
		if claims.Authorised == false {
			return errForbidden("2fa required")
		}
		return next(c)
	}
}
