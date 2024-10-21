package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewSpell(db *gorm.DB, s *models.Spell) {
	fillSpellDetails(s)
	db.Create(s)
}

func fillSpellDetails(s *models.Spell) {
	if s.Name == "" {
		s.Name = random.String(16)
	}
}
