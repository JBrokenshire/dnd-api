package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetBackgrounds() {
	backgrounds := []*m.Background{
		{
			ID:          1,
			Name:        "Nomadic Troubadour",
			Feature:     "Song of Hospitality",
			Description: "Wherever you wander, people are drawn to your performances. In most civilized places, you can secure food, shelter, or safe passage for yourself (and sometimes your companions) by offering music, storytelling, or other entertainment. Even in harsher wilds, your charm can open doors: a campfire song among strangers, a story told to fellow travelers, or a rhythm shared with other nomads. While this rarely earns you coin, it ensures that you are rarely without a place to rest.",
		},
	}

	for _, background := range backgrounds {
		err := s.DB.Where("id = ?", background.ID).FirstOrCreate(&background).Error
		if err != nil {
			log.Printf("Error create background with id %v in seeder: %v", background.ID, err.Error())
		}
	}
}
