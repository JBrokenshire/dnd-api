package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"purplevisits.com/mdm/db/models"
	"purplevisits.com/mdm/sso/session"
	"strings"
)

// LoadSsoOwner Follows on from validating JWT and pulls the owner out and will add to context.
func LoadSsoOwner(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			owner := c.Get("user").(*jwt.Token)

			claims := owner.Claims.(*session.JwtSSOClaims)
			uid := claims.Subject
			enterpriseId := claims.EnterpriseId

			// Pull the owner and make sure they have the enterprise_uid set
			if uid == "" {
				return errUnauthorized("Invalid OwnerId in request")
			}
			if enterpriseId == "" {
				return errUnauthorized("Invalid Enterprise in request")
			}

			// The owner UID should be either a UUID or the enterpriseID + "-" + NOMIS ID for BT integrations
			if !strings.HasPrefix(uid, enterpriseId) && len(uid) < 10 {
				log.Printf("ERROR: (LoadSsoOwner) A token has been received where the UID does not start with the enterpirse ID and the UID is less than 10 characters. Enterprise: %v   UID: %v", enterpriseId, uid)
				uid = enterpriseId + "-" + uid
			}

			// Pull the owner out of the database
			var currentOwner models.Owner
			err := db.Where("uid = ?", uid).First(&currentOwner).Error
			if err != nil {
				log.Printf("(LoadSsoOwner) Error finding owner: %v   UID: %v", err, uid)
				return errUnauthorized("Invalid OwnerId in request")
			}

			// We can now make sure this owner is part of the enterprise in the JWT. Not required but just as an extra
			// check
			if currentOwner.EnterpriseUid != enterpriseId {
				return errUnauthorized("Incorrect Enterprise in request")
			}

			c.Set("enterpriseId", enterpriseId)
			c.Set("uid", uid)
			c.Set("currentOwner", &currentOwner)
			return next(c)
		}
	}
}
