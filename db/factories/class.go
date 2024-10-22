package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewClass(db *gorm.DB, c *models.Class) {
	fillClassDetails(c)
	db.Create(c)
}

func fillClassDetails(c *models.Class) {
	if c.Name == "" {
		c.Name = random.String(16)
	}
}
