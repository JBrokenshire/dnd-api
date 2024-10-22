package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewSubclass(db *gorm.DB, s *models.Subclass) {
	fillSubclassDetails(s)
	db.Create(s)
}

func fillSubclassDetails(s *models.Subclass) {
	if s.ClassID == 0 {
		s.ClassID = 1
	}
	if s.Name == "" {
		s.Name = random.String(16)
	}
}
