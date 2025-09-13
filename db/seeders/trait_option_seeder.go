package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetTraitOptions() {
	traitOptions := []*m.TraitOption{
		{
			ID:          1,
			TraitId:     "dragonborn-draconic-ancestry",
			Name:        "Brass Dragon",
			Description: "Fire",
		},
		{
			ID:          2,
			TraitId:     "dragonborn-breath-weapon",
			Name:        "Brass Dragon",
			Description: "As an action once per short rest, exhale a 5 by 30ft. line (DEX DC 16, half damage on success) for 2d6 Fire Damage [6th] 3d6, [11th] 4d6, [16th] 5d6",
		},
		{
			ID:          3,
			TraitId:     "dragonborn-damage-resistance",
			Name:        "Brass Dragon",
			Description: "Fire Damage",
		},
	}

	for _, traitOption := range traitOptions {
		err := s.DB.Where("id = ?", traitOption.ID).FirstOrCreate(&traitOption).Error
		if err != nil {
			log.Printf("Error creating trait option with id %v in seeder: %v", traitOption.ID, err.Error())
		}
	}
}
