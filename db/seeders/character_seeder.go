package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacters() {
	characters := []m.Character{
		{
			ID:               1,
			UserId:           1,
			Name:             "Axel Claystride",
			ClassId:          1,
			RaceId:           1,
			Level:            6,
			Pronouns:         "He/Him",
			BackgroundId:     1,
			Alignment:        "Chaotic Good",
			Gender:           "Male",
			Eyes:             "Dark Brown",
			Size:             "Medium",
			Height:           "6' 0\"",
			Hair:             "Black, Short and Spiky",
			Skin:             "Rocky, Gray and Cracked",
			Age:              "23",
			Weight:           "200lb.",
			Strength:         9,
			Dexterity:        16,
			Constitution:     14,
			Intelligence:     11,
			Wisdom:           12,
			Charisma:         16,
			CurrentHitPoints: 38,
			MaxHitPoints:     38,
			AdvancementType:  m.AdvancementTypeMilestone,
			HitPointType:     m.HitPointTypeManual,
			Proficiencies:    `{"armour":["Light Armour"], "weapons":["Simple Weapons"], "tools":["Lute","Zulkoon","Painters' Supplies"], "languages":["Common","Primordial","Celestial"}`,
		},
	}

	for _, character := range characters {
		err := s.DB.Where("id = ?", character.ID).FirstOrCreate(&character).Error
		if err != nil {
			log.Printf("Error creating character with id %v in seeder: %v", character.ID, err.Error())
		}
	}
}
