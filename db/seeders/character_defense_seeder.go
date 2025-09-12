package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterDefenses() {
	characterDefenses := []m.CharacterDefense{
		{
			ID:          1,
			CharacterID: 2,
			DefenseType: m.DefenseTypeResistance,
			DamageType:  m.DamageTypeFire,
		},
	}

	for _, defense := range characterDefenses {
		err := s.DB.Where("id = ?", defense.ID).FirstOrCreate(&defense).Error
		if err != nil {
			log.Printf("Error creating character defense with id %v in seeder: %v", defense.ID, err.Error())
		}
	}
}
