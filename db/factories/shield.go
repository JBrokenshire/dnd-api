package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewShield(db *gorm.DB, s *models.Shield) {
	fillShieldDetails(s)
	db.Create(s)
}

func fillShieldDetails(s *models.Shield) {
	if s.BonusAC == 0 {
		s.BonusAC = 2
	}
}
