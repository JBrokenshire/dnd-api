package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"purplevisits.com/mdm/db/models"
	"purplevisits.com/mdm/pkg/jwt_service"
	"time"
)

// LoadUser Follows on from validating JWT and pulls the user out and will add to context. Also adds users
// permissions in the object so that we can validate permissions added to routes.
func LoadUser(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*jwt_service.JwtCustomClaims)
			uid := claims.ID
			enterpriseId := claims.EnterpriseId

			// Pull the user and make sure they have the enterprise_uid set
			if uid == "" {
				return errForbidden("Missing UserId in request")
			}
			if enterpriseId == "" {
				return errForbidden("Missing Enterprise in request")
			}

			// Pull the user out of the database
			var currentUser models.User
			err := db.Preload("Roles.Permissions").Where("uid = ?", uid).First(&currentUser).Error
			if err != nil {
				log.Printf("Error finding user: %v", err)
				return errForbidden("Invalid UserId in request")
			}

			// We can now make sure this user is part of the enterprise in the JWT. Not required but just as an extra
			// check. Exclude super admins
			if currentUser.SuperAdmin == false && currentUser.EnterpriseUID != enterpriseId {
				return errForbidden("Incorrect Enterprise in request")
			}

			// Update last seen if previous last seen was more than 5 minutes ago
			if currentUser.LastSeen == nil || currentUser.LastSeen.Add(5*time.Minute).Before(time.Now()) {
				err := db.Model(&models.User{}).Where("id = ?", currentUser.ID).Update("last_seen", time.Now()).Error
				if err != nil {
					log.Printf("Error updating user last seen: %v", err)
				}
			}

			// Rather than adding all roles, we should add all the permissions that this user has to their user object.
			// This will let us easily see if they have a permission, regardless of what role it was attached to.
			c.Set("enterpriseId", enterpriseId)
			c.Set("uid", uid)
			c.Set("currentUser", &currentUser)
			return next(c)
		}
	}
}
