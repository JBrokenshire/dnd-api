package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"purplevisits.com/mdm/db/models"
	"purplevisits.com/mdm/pkg/jwt_service"
	"strings"
	"time"
)

// LoadOwner Follows on from validating JWT and pulls the owner out and will add to context.
func LoadOwner(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			owner := c.Get("user").(*jwt.Token)

			claims := owner.Claims.(*jwt_service.JwtOwnerCustomClaims)
			uid := claims.ID
			enterpriseId := claims.EnterpriseId
			packages := claims.Packages
			deviceUid := claims.DeviceUid

			// Pull the owner and make sure they have the enterprise_uid set
			if uid == "" {
				return errForbidden("Invalid OwnerId in request")
			}
			if enterpriseId == "" {
				return errForbidden("Invalid Enterprise in request")
			}
			if len(packages) < 1 {
				return errForbidden("Invalid packages in request")
			}

			if claims.Type != jwt_service.TokenTypeOwner {
				return errForbidden("Invalid token type. Expecting Owner")
			}

			// The owner UID should be either a UUID or the enterpriseID + "-" + NOMIS ID for BT integrations
			if !strings.HasPrefix(uid, enterpriseId) && len(uid) < 10 {
				log.Printf("ERROR: (LoadOwner) A token has been received where the UID does not start with the enterpirse ID and the UID is less than 10 characters. Enterprise: %v   UID: %v", enterpriseId, uid)
				uid = enterpriseId + "-" + uid
			}

			// Pull the owner out of the database
			var currentOwner models.Owner
			err := db.Where("uid = ?", uid).First(&currentOwner).Error
			if err != nil {
				log.Printf("(LoadOwner) Error finding owner: %v   OwnerID: %v", err, uid)
				return errForbidden("Invalid OwnerId in request")
			}

			// We can now make sure this owner is part of the enterprise in the JWT. Not required but just as an extra
			// check
			if currentOwner.EnterpriseUid != enterpriseId {
				return errForbidden("Incorrect Enterprise in request")
			}

			// Update last seen if previous last seen was more than 5 minutes ago
			if currentOwner.LastSeen == nil || currentOwner.LastSeen.Add(5*time.Minute).Before(time.Now()) {
				err := db.Model(&models.Owner{}).Where("id = ?", currentOwner.ID).Update("last_seen", time.Now()).Error
				if err != nil {
					log.Printf("Error updating owner last seen: %v", err)
				}
			}

			c.Set("enterpriseId", enterpriseId)
			c.Set("uid", uid)
			c.Set("currentOwner", &currentOwner)
			c.Set("packages", packages)
			c.Set("deviceUid", deviceUid)
			return next(c)
		}
	}
}

// IsApp checks that the owner token is from a specific application
func IsApp(appId string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenPackages := c.Get("packages").([]string)
			if len(tokenPackages) < 1 {
				return errForbidden("Missing Application Id. This application is not authenticated for this endpoint")
			}
			for _, app := range tokenPackages {
				if app == appId {
					return next(c)
				}
			}
			return errForbidden("Invalid Application Id. This application is not authenticated for this endpoint")
		}
	}
}
