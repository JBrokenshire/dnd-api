package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetBackgrounds() {
	backgrounds := []m.Background{
		{
			ID:          1,
			Name:        "Nomadic Troubadour",
			Feature:     "Song of Hospitality",
			Description: "Something about getting a free bed and food in exchange for entertainment",
		},
	}

	for _, background := range backgrounds {
		err := s.DB.Where("id = ?", background.ID).FirstOrCreate(&background).Error
		if err != nil {
			log.Printf("Error creating background with id %v in seeder: %v", background.ID, err.Error())
		}
	}
}
