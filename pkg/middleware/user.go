package middleware

import (
	"dnd-api/db/models"
	"dnd-api/services/jwt_service"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

// LoadUser Follows on from validating JWT and pulls the user out and will add to context. Also adds users
// permissions in the object so that we can validate permissions added to routes.
func LoadUser(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*jwt_service.JwtCustomClaims)
			id := claims.ID

			if id == 0 {
				return errForbidden("Missing UserId in request")
			}

			// Pull the user out of the database
			var currentUser models.User
			err := db.Where("id = ?", id).First(&currentUser).Error
			if err != nil {
				log.Printf("Error finding user: %v", err)
				return errForbidden("Invalid UserId in request")
			}

			// Update last seen if previous last seen was more than 5 minutes ago
			if currentUser.LastSeen == nil || currentUser.LastSeen.Add(5*time.Minute).Before(time.Now()) {
				err := db.Model(&models.User{}).Where("id = ?", currentUser.ID).Update("last_seen", time.Now()).Error
				if err != nil {
					log.Printf("Error updating user last seen: %v", err)
				}
			}

			c.Set("currentUser", &currentUser)
			return next(c)
		}
	}
}
