package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetRaceTraits() {
	raceTraits := []*m.RaceTrait{
		{
			ID:      1,
			RaceId:  1,
			TraitId: "earth-genasi-asi",
		},
		{
			ID:      2,
			RaceId:  1,
			TraitId: "earth-genasi-earth-walk",
		},
		{
			ID:      3,
			RaceId:  1,
			TraitId: "earth-genasi-merge-with-stone",
		},
		{
			ID:      4,
			RaceId:  2,
			TraitId: "dragonborn-asi",
		},
		{
			ID:      5,
			RaceId:  2,
			TraitId: "dragonborn-draconic-ancestry",
		},
		{
			ID:      6,
			RaceId:  2,
			TraitId: "dragonborn-breath-weapon",
		},
		{
			ID:      7,
			RaceId:  2,
			TraitId: "dragonborn-damage-resistance",
		},
		{
			ID:      8,
			RaceId:  3,
			TraitId: "stout-halfling-asi",
		},
		{
			ID:      9,
			RaceId:  3,
			TraitId: "stout-halfling-lucky",
		},
		{
			ID:      10,
			RaceId:  3,
			TraitId: "stout-halfling-brave",
		},
		{
			ID:      11,
			RaceId:  3,
			TraitId: "stout-halfling-halfling-nimbleness",
		},
		{
			ID:      12,
			RaceId:  3,
			TraitId: "stout-halfling-stout-resilience",
		},
	}

	for _, raceTrait := range raceTraits {
		err := s.DB.Where("id = ?", raceTrait.ID).FirstOrCreate(&raceTrait).Error
		if err != nil {
			log.Printf("Error creating race trait with id %v in seeder: %v", raceTrait.ID, err.Error())
		}
	}
}
