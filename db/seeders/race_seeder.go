package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetRaces() {
	races := []m.Race{
		{
			ID:               1,
			Name:             "Earth Genasi",
			ShortDescription: "Earth Genasi",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 5-6 feet tall)",
			BaseSpeed:        30,
		},
	}

	for _, race := range races {
		err := s.DB.Where("id = ?", race.ID).FirstOrCreate(&race).Error
		if err != nil {
			log.Printf("Error creating race with id %v in seeder: %v", race.ID, err.Error())
		}
	}
}
