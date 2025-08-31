package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterSpells() {
	characterSpells := []*m.CharacterSpell{
		{
			ID:          1,
			CharacterID: 1,
			SpellID:     1,
		},
		{
			ID:          2,
			CharacterID: 1,
			SpellID:     2,
		},
		{
			ID:          3,
			CharacterID: 1,
			SpellID:     3,
		},
		{
			ID:          4,
			CharacterID: 1,
			SpellID:     4,
		},
		{
			ID:          5,
			CharacterID: 1,
			SpellID:     5,
		},
		{
			ID:          6,
			CharacterID: 1,
			SpellID:     6,
		},
		{
			ID:          7,
			CharacterID: 1,
			SpellID:     7,
		},
		{
			ID:          8,
			CharacterID: 1,
			SpellID:     8,
		},
		{
			ID:          9,
			CharacterID: 1,
			SpellID:     9,
		},
		{
			ID:          10,
			CharacterID: 1,
			SpellID:     10,
		},
		{
			ID:          11,
			CharacterID: 1,
			SpellID:     11,
		},
		{
			ID:          12,
			CharacterID: 1,
			SpellID:     12,
		},
		{
			ID:          13,
			CharacterID: 1,
			SpellID:     13,
		},
		{
			ID:          14,
			CharacterID: 1,
			SpellID:     14,
		},
	}

	for _, spell := range characterSpells {
		err := s.DB.Where("id = ?", spell.ID).FirstOrCreate(&spell).Error
		if err != nil {
			log.Printf("Error creating spell with id %v in seeder: %v", spell.ID, err.Error())
		}
	}
}
