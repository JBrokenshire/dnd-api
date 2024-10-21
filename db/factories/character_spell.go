package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterSpell(db *gorm.DB, cs *models.CharacterSpell) {
	fillCharacterSpellDetails(cs)
	db.Create(cs)
}

func fillCharacterSpellDetails(cs *models.CharacterSpell) {
	if cs.Origin == "" {
		cs.Origin = "Wizard"
	}
}
