package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewSpell(db *gorm.DB, spell *m.Spell) {
	fillSpellDefaults(spell)
	err := db.Create(spell).Error
	if err != nil {
		log.Println("Error creating spell in factory: ", err.Error())
	}
}

func fillSpellDefaults(s *m.Spell) {
	if s.Name == "" {
		s.Name = random.String(16)
	}
	if s.School == "" {
		s.School = m.MagicSchoolAbjuration
	}
	if s.CastingTime == "" {
		s.CastingTime = "1A"
	}
	if s.Distance == "" {
		s.Distance = random.String(8)
	}
	if s.Notes == "" {
		s.Notes = random.String(16)
	}
}
