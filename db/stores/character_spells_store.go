package stores

import (
	"dnd-api/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"sort"
)

type CharacterSpellsStore interface {
	GetHasSpellsByCharacterID(id interface{}) (*bool, error)
	GetSpellsByCharacterID(id interface{}) ([]*models.CharacterSpell, error)
}

type GormCharacterSpellsStore struct {
	DB *gorm.DB
}

func NewGormCharacterSpellsStore(db *gorm.DB) *GormCharacterSpellsStore {
	return &GormCharacterSpellsStore{
		DB: db,
	}
}

func (g *GormCharacterSpellsStore) GetHasSpellsByCharacterID(id interface{}) (*bool, error) {
	var spell models.CharacterSpell
	g.DB.Table("character_spells").Where("character_id = ?", id).First(&spell)

	hasSpells := spell.ID != 0
	fmt.Println(">>> ", spell.ID)

	return &hasSpells, nil
}

func (g *GormCharacterSpellsStore) GetSpellsByCharacterID(id interface{}) ([]*models.CharacterSpell, error) {
	var spells []*models.CharacterSpell
	err := g.DB.Preload("Spell").Where("character_id = ?", id).Find(&spells).Error
	if err != nil {
		return nil, err
	}

	sort.Slice(spells, func(i, j int) bool {
		return spells[1].Spell.Level < spells[j].Spell.Level
	})

	return spells, nil
}
