package middleware

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"purplevisits.com/mdm/db/models"
	"purplevisits.com/mdm/pkg/ip"
)

// OwnerIpCheck check that the IP address is within enterprise.ownerIpRanges setting
func OwnerIpCheck(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			enterpriseUid := c.Get("enterpriseId").(string)

			// Pull the enterprise out of the database
			var enterprise models.Enterprise
			err := db.Where("uid = ?", enterpriseUid).First(&enterprise).Error
			if err != nil {
				log.Printf("Error finding enterprise: %v", err)
				return errForbidden("Invalid EnterpriseUid in request")
			}

			// if enterprise has ip range set for owners then check ip is within the range
			isValidIp := ip.InRanges(c.RealIP(), enterprise.OwnerIpRanges)
			if !isValidIp {
				return errForbidden("Owner Ip address is not in enterprise range")
			}

			return next(c)
		}
	}
}
