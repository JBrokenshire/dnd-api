package repositories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

type CharacterUsedSpellSlotRepository struct {
	*Repository
}

func NewCharacterUsedSpellSlotRepository(db *gorm.DB) *CharacterUsedSpellSlotRepository {
	return &CharacterUsedSpellSlotRepository{
		&Repository{
			Db: db,
		},
	}
}

func (r *CharacterUsedSpellSlotRepository) Get(characterID, spellLevel interface{}) *m.CharacterUsedSpellSlot {
	var usedSpellSlot m.CharacterUsedSpellSlot
	r.Db.
		Where("character_id = ?", characterID).
		Where("spell_level = ?", spellLevel).
		First(&usedSpellSlot)
	return &usedSpellSlot
}
