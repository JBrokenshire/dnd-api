package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacters() {
	characters := []models.Character{
		{
			ID:                     1,
			Name:                   "Faelan Haversham",
			ProfilePictureURL:      "https://www.dndbeyond.com/avatars/40171/941/1581111423-117906084.jpeg?width=150&height=150&fit=crop&quality=95&auto=webp",
			Level:                  3,
			ClassID:                3,
			RaceID:                 18,
			Strength:               12,
			Dexterity:              16,
			Constitution:           14,
			Intelligence:           10,
			Wisdom:                 15,
			Charisma:               8,
			CurrentHitPoints:       28,
			MaxHitPoints:           28,
			ProficientDexterity:    true,
			ProficientIntelligence: true,
			AttacksPerAction:       1,
		},
		{
			ID:                     2,
			Name:                   "PeeWee McAnkle-Biter",
			Level:                  5,
			ProfilePictureURL:      "https://www.dndbeyond.com/avatars/17/969/636378926856690034.png?width=150&height=150&fit=crop&quality=95&auto=webp",
			ClassID:                1,
			RaceID:                 3,
			Strength:               16,
			Dexterity:              14,
			Constitution:           16,
			Intelligence:           12,
			Wisdom:                 11,
			Charisma:               8,
			WalkingSpeedModifier:   10,
			CurrentHitPoints:       49,
			MaxHitPoints:           60,
			ProficientStrength:     true,
			ProficientConstitution: true,
			InitiativeModifier:     5,
			AttacksPerAction:       2,
		},
		{
			ID:                     3,
			Name:                   "Iratham Veomakute",
			Level:                  3,
			ClassID:                3,
			RaceID:                 14,
			ProfilePictureURL:      "https://www.dndbeyond.com/avatars/36645/678/1581111423-94761552.jpeg?width=150&height=150&fit=crop&quality=95&auto=webp",
			Strength:               18,
			Dexterity:              13,
			Constitution:           14,
			Intelligence:           15,
			Wisdom:                 11,
			Charisma:               11,
			CurrentHitPoints:       34,
			MaxHitPoints:           34,
			ProficientDexterity:    true,
			ProficientIntelligence: true,
			AttacksPerAction:       1,
		},
		{
			ID:                     4,
			Name:                   "Kael Drakeshield",
			Level:                  4,
			ClassID:                6,
			RaceID:                 2,
			ProfilePictureURL:      "https://www.dndbeyond.com/avatars/42817/335/1581111423-124892113.jpeg?width=150&height=150&fit=crop&quality=95&auto=webp",
			Strength:               19,
			Dexterity:              15,
			Constitution:           19,
			Intelligence:           6,
			Wisdom:                 14,
			Charisma:               11,
			Inspiration:            true,
			CurrentHitPoints:       43,
			MaxHitPoints:           43,
			ProficientStrength:     true,
			ProficientConstitution: true,
			AttacksPerAction:       1,
		},
	}

	for _, character := range characters {
		err := s.DB.Where("id = ?", character.ID).FirstOrCreate(&character).Error
		if err != nil {
			log.Printf("error creating character %s: %s", character.Name, err.Error())
		}
	}
}
