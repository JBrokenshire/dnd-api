package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetSubclasses() {
	subclasses := []models.Subclass{
		{
			ID:      1,
			ClassID: 1,
			Name:    "Path of the Totem Warrior",
		},
		{
			ID:      2,
			ClassID: 3,
			Name:    "Order of the Mutant",
		},
		{
			ID:      3,
			ClassID: 3,
			Name:    "Order of the Lycan",
		},
		{
			ID:      4,
			ClassID: 6,
			Name:    "Echo Knight",
		},
		{
			ID:      5,
			ClassID: 11,
			Name:    "Lunar Sorcery",
		},
	}

	for _, subclass := range subclasses {
		err := s.DB.Where("id = ?", subclass.ID).FirstOrCreate(&subclass).Error
		if err != nil {
			log.Printf("error creating subclass with id %v - %v", subclass.ID, err)
		}
	}
}
