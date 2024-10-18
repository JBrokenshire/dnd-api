package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterSpells() {
	characterSpells := []models.CharacterSpell{
		{
			ID:          1,
			CharacterID: 5,
			SpellID:     1,
		},
		{
			ID:          2,
			CharacterID: 5,
			SpellID:     2,
		},
		{
			ID:          3,
			CharacterID: 5,
			SpellID:     3,
		},
		{
			ID:          4,
			CharacterID: 5,
			SpellID:     4,
		},
		{
			ID:          5,
			CharacterID: 5,
			SpellID:     5,
		},
		{
			ID:          6,
			CharacterID: 5,
			SpellID:     7,
		},
		{
			ID:          7,
			CharacterID: 5,
			SpellID:     8,
		},
		{
			ID:          8,
			CharacterID: 5,
			SpellID:     9,
		},
		{
			ID:          9,
			CharacterID: 5,
			SpellID:     10,
		},
		{
			ID:          10,
			CharacterID: 5,
			SpellID:     11,
		},
		{
			ID:          11,
			CharacterID: 5,
			SpellID:     12,
		},
		{
			ID:          12,
			CharacterID: 5,
			SpellID:     13,
		},
		{
			ID:          13,
			CharacterID: 4,
			SpellID:     6,
		},
	}

	for _, characterSpell := range characterSpells {
		err := s.DB.Where("id = ?", characterSpell.ID).FirstOrCreate(&characterSpell).Error
		if err != nil {
			log.Printf("error creating character spell for character '%v', spell '%v' - %v", characterSpell.CharacterID, characterSpell.SpellID, err)
		}
	}
}
