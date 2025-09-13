package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

func NewCharacterUsedSpellSlot(db *gorm.DB, usedSpellSlot *m.CharacterUsedSpellSlot) {
	fillCharacterUsedSpellSlotsDefaults(usedSpellSlot)
	err := db.Create(usedSpellSlot).Error
	if err != nil {
		log.Println("Error creating character used spell slot in factory: ", err)
	}
}

func fillCharacterUsedSpellSlotsDefaults(usedSpellSlot *m.CharacterUsedSpellSlot) {
	if usedSpellSlot.SpellLevel == 0 {
		usedSpellSlot.SpellLevel = 1
	}
}
