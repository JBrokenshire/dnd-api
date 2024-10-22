package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetRaceTraits() {
	raceTraits := []models.RaceTrait{
		{
			ID:      1,
			RaceID:  18,
			TraitID: 1,
		},
		{
			ID:      2,
			RaceID:  18,
			TraitID: 2,
		},
		{
			ID:      3,
			RaceID:  18,
			TraitID: 3,
		},
		{
			ID:      4,
			RaceID:  18,
			TraitID: 4,
		},
		{
			ID:      5,
			RaceID:  3,
			TraitID: 5,
		},
		{
			ID:      6,
			RaceID:  3,
			TraitID: 6,
		},
		{
			ID:      7,
			RaceID:  3,
			TraitID: 7,
		},
		{
			ID:      8,
			RaceID:  3,
			TraitID: 8,
		},
		{
			ID:      9,
			RaceID:  3,
			TraitID: 9,
		},
		{
			ID:      10,
			RaceID:  14,
			TraitID: 10,
		},
		{
			ID:      11,
			RaceID:  14,
			TraitID: 11,
		},
		{
			ID:      12,
			RaceID:  14,
			TraitID: 12,
		},
		{
			ID:      13,
			RaceID:  14,
			TraitID: 13,
		},
		{
			ID:      14,
			RaceID:  2,
			TraitID: 14,
		},
		{
			ID:      15,
			RaceID:  2,
			TraitID: 15,
		},
		{
			ID:      16,
			RaceID:  2,
			TraitID: 16,
		},
		{
			ID:      17,
			RaceID:  2,
			TraitID: 17,
		},
		{
			ID:      18,
			RaceID:  17,
			TraitID: 18,
		},
		{
			ID:      19,
			RaceID:  17,
			TraitID: 1,
		},
		{
			ID:      20,
			RaceID:  17,
			TraitID: 2,
		},
		{
			ID:      21,
			RaceID:  17,
			TraitID: 3,
		},
		{
			ID:      24,
			RaceID:  17,
			TraitID: 19,
		},
		{
			ID:      25,
			RaceID:  17,
			TraitID: 20,
		},
	}

	for _, raceTrait := range raceTraits {
		err := s.DB.Where("id = ?", raceTrait.ID).FirstOrCreate(&raceTrait).Error
		if err != nil {
			log.Printf("error create race trait with id %v - %v", raceTrait.ID, err)
		}
	}
}
