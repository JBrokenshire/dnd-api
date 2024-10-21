package stores

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"sort"
)

type CharacterSpellsStore interface {
	GetHasSpellsByCharacterID(id interface{}) (*bool, error)
	GetSpellsByCharacterID(id interface{}) ([]*models.CharacterSpell, error)
	GetAttackSpellsByCharacterID(id interface{}) ([]*models.CharacterSpell, error)
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

	return &hasSpells, nil
}

func (g *GormCharacterSpellsStore) GetSpellsByCharacterID(id interface{}) ([]*models.CharacterSpell, error) {
	var spells []*models.CharacterSpell
	err := g.DB.Preload("Spell").Where("character_id = ?", id).Find(&spells).Error
	if err != nil {
		return nil, err
	}

	sort.Slice(spells, func(i, j int) bool {
		return spells[i].Spell.Level < spells[j].Spell.Level
	})

	return spells, nil
}

func (g *GormCharacterSpellsStore) GetAttackSpellsByCharacterID(id interface{}) ([]*models.CharacterSpell, error) {
	var characterSpells []*models.CharacterSpell
	err := g.DB.
		Preload("Spell", "effect = ? AND effect IS NOT NULL", "Attack").
		Where("character_id = ?", id).
		Find(&characterSpells).Error
	if err != nil {
		return nil, err
	}

	var attackSpells []*models.CharacterSpell
	for _, characterSpell := range characterSpells {
		if characterSpell.Spell.ID == 0 {
			continue
		}

		attackSpells = append(attackSpells, characterSpell)
	}

	return attackSpells, nil
}
