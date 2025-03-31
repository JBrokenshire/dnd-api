package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"purplevisits.com/mdm/db/models"
	"purplevisits.com/mdm/pkg/jwt_service"
)

// CheckEnrollmentToken Follows on from validating JWT and checked it's a valid enrollment token
func CheckEnrollmentToken(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			deviceToken := c.Get("user").(*jwt.Token)
			claims := deviceToken.Claims.(*jwt_service.JwtCustomClaims)
			deviceUid := claims.ID
			enterpriseId := claims.EnterpriseId

			// Pull the owner and make sure they have the enterprise_uid set
			if deviceUid == "" {
				return errForbidden("Invalid deviceId in request")
			}
			if enterpriseId == "" {
				return errForbidden("Invalid Enterprise in request")
			}

			if claims.Type != jwt_service.TokenTypeDevice {
				return errForbidden("Invalid token type. Expecting Device")
			}

			// Check the device belongs to this enterprise
			var device models.Device
			err := db.Unscoped().Where("uid = ? and enterprise_uid = ?", deviceUid, enterpriseId).First(&device).Error
			if err != nil {
				return errForbidden(fmt.Sprintf("Invalid Device in request, device uid: %v, enterprise_uid: %v", deviceUid, enterpriseId))
			}

			ownerUid := claims.OwnerUid

			c.Set("enterpriseId", enterpriseId)
			c.Set("deviceUid", deviceUid)
			c.Set("ownerUid", ownerUid)
			return next(c)
		}
	}
}
