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
			Origin:      "Sorcerer",
		},
		{
			ID:          2,
			CharacterID: 5,
			SpellID:     2,
			Origin:      "Sorcerer",
		},
		{
			ID:          3,
			CharacterID: 5,
			SpellID:     3,
			Origin:      "Sorcerer",
		},
		{
			ID:          4,
			CharacterID: 5,
			SpellID:     4,
			Origin:      "Sorcerer",
		},
		{
			ID:          5,
			CharacterID: 5,
			SpellID:     5,
			Origin:      "Sorcerer",
		},
		{
			ID:          6,
			CharacterID: 5,
			SpellID:     7,
			Origin:      "Sorcerer",
		},
		{
			ID:          7,
			CharacterID: 5,
			SpellID:     8,
			Origin:      "Sorcerer",
		},
		{
			ID:          8,
			CharacterID: 5,
			SpellID:     9,
			Origin:      "Sorcerer",
		},
		{
			ID:          9,
			CharacterID: 5,
			SpellID:     10,
			Origin:      "Sorcerer",
		},
		{
			ID:          10,
			CharacterID: 5,
			SpellID:     11,
			Origin:      "Sorcerer",
		},
		{
			ID:          11,
			CharacterID: 5,
			SpellID:     12,
			Origin:      "Sorcerer",
		},
		{
			ID:          12,
			CharacterID: 5,
			SpellID:     13,
			Origin:      "Sorcerer",
		},
		{
			ID:          13,
			CharacterID: 4,
			SpellID:     6,
			Origin:      "Aberrant Dragonmark",
		},
		{
			ID:          14,
			CharacterID: 4,
			SpellID:     14,
			Origin:      "Aberrant Dragonmark",
		},
	}

	for _, characterSpell := range characterSpells {
		err := s.DB.Where("id = ?", characterSpell.ID).FirstOrCreate(&characterSpell).Error
		if err != nil {
			log.Printf("error creating character spell for character '%v', spell '%v' - %v", characterSpell.CharacterID, characterSpell.SpellID, err)
		}
	}
}
