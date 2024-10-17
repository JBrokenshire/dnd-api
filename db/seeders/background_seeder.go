package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetBackgrounds() {
	backgrounds := []models.Background{
		{
			Name:        "Outlander",
			Feature:     "Wanderer",
			Description: "You have an excellent memory for maps and geography, and you can always recall the general layout of terrain, settlements, and other features around you. In addition, you can find food and fresh water for yourself and up to five other people each day, provided that the land offers berries, small game, water, and so forth.",
		},
		{
			Name:        "Soldier",
			Feature:     "Military Rank",
			Description: "You have a military rank from your career as a soldier. Soldiers loyal to your former military organization still recognize your authority and influence, and they defer to you if they are of a lower rank. You can invoke your rank to exert influence over other soldiers and requisition simple equipment or horses for temporary use. You can also usually gain access to friendly military encampments and fortresses where your rank is recognized.",
		},
		{
			Name:        "Urban Bounty Hunter",
			Feature:     "Ear to the Ground",
			Description: "You are in frequent contact with people in the segment of society that your chosen quarries move through. These people might be associated with the criminal underworld, the rough-and-tumble folk of the streets, or members of high society. This connection comes in the form of a contact in any city you visit, a person who provides information about the people and places of the local area.",
		},
	}

	for _, background := range backgrounds {
		err := s.DB.Where("name = ?", background.Name).FirstOrCreate(&background).Error
		if err != nil {
			log.Printf("error create background: %v - %v", background.Name, err)
		}
	}
}
