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
			Description: "Wherever you wander, people are drawn to your performances. In most civilized places, you can secure food, shelter, or safe passage for yourself (and sometimes your companions) by offering music, storytelling, or other entertainment. Even in harsher wilds, your charm can open doors: a campfire song among strangers, a story told to fellow travelers, or a rhythm shared with other nomads. While this rarely earns you coin, it ensures that you are rarely without a place to rest.",
		},
		{
			ID:          2,
			Name:        "Urban Bounty Hunter",
			Feature:     "Ear to the Ground",
			Description: "You are in frequent contact with people in the segment of society that your chosen quarries move through. These people might be associated with the criminal underworld, the rough-and-tumble folk of the streets, or members of high society. This connection comes in the form of a contact in any city you visit, a person who provides information about the people and places of the local area.",
		},
		{
			ID:          3,
			Name:        "Outlander",
			Feature:     "Wanderer",
			Description: "You have an excellent memory for maps and geography, and you can always recall the general layout of terrain, settlements, and other features around you. In addition, you can find food and fresh water for yourself and up to five other people each day, provided that the land offers berries, small game, water, and so forth.",
		},
	}

	for _, background := range backgrounds {
		err := s.DB.Where("id = ?", background.ID).FirstOrCreate(&background).Error
		if err != nil {
			log.Printf("Error creating background with id %v in seeder: %v", background.ID, err.Error())
		}
	}
}
