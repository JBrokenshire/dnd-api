package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterSelectedRaceTraitOptions() {
	characterSelectedRaceTraitOptions := []*m.CharacterSelectedRaceTraitOption{
		{
			ID:            1,
			CharacterId:   2,
			TraitOptionId: 1,
		},
		{
			ID:            2,
			CharacterId:   2,
			TraitOptionId: 2,
		},
		{
			ID:            3,
			CharacterId:   2,
			TraitOptionId: 3,
		},
	}

	for _, characterSelectedRaceTraitOption := range characterSelectedRaceTraitOptions {
		err := s.DB.Where("id = ?", characterSelectedRaceTraitOption.ID).FirstOrCreate(&characterSelectedRaceTraitOption).Error
		if err != nil {
			log.Printf("Error creating character selected race trait option with id %v in seeder: %v", characterSelectedRaceTraitOption.ID, err.Error())
		}
	}
}
