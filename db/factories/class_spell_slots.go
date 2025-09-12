package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

func NewClassSpellSlots(db *gorm.DB, spellSlots *m.ClassSpellSlots) {
	fillClassSpellSlotsDefaults(spellSlots)
	err := db.Create(spellSlots).Error
	if err != nil {
		log.Println("Error create class spell slots in factory: ", err.Error())
	}
}

func fillClassSpellSlotsDefaults(s *m.ClassSpellSlots) {
	if s.ClassLevel == 0 {
		s.ClassLevel = 1
	}
	if s.SpellLevel == 0 {
		s.SpellLevel = 2
	}
	if s.SpellSlots == 0 {
		s.SpellSlots = 1
	}
}
