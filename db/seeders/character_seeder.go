package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacters() {
	characters := []*m.Character{
		{
			ID:               1,
			UserId:           1,
			Name:             "Axel Claystride",
			ClassId:          2,
			RaceId:           12,
			Level:            6,
			Pronouns:         "He/Him",
			BackgroundId:     1,
			Alignment:        "Chaotic Good",
			Gender:           "Male",
			Eyes:             "Dark Brown",
			Size:             "Medium",
			Height:           `6'0"'`,
			Hair:             "Black, short & spiky",
			Skin:             "Grey, rocky & cracked",
			Age:              "23",
			Weight:           "200 lb.",
			Strength:         9,
			Dexterity:        16,
			Constitution:     14,
			Intelligence:     11,
			Wisdom:           12,
			Charisma:         18,
			CurrentHitPoints: 38,
			MaxHitPoints:     38,
			AdvancementType:  "Milestone",
			HitPointType:     "Manual",
			Senses:           "[]",
			Proficiencies:    `{"Armour": ["Light Armour"], "Weapons": ["Simple Weapons"], "Tools": ["Drum", "Lute", "Zulkoon", "Painter's Supplies"], "Languages": ["Common", "Elvish", "Primordial"]}`,
		},
	}

	for _, character := range characters {
		err := s.DB.Where("id = ?", character.ID).FirstOrCreate(&character).Error
		if err != nil {
			log.Printf("Error creating character with id %v in seeder: %v", character.ID, err.Error())
		}
	}
}
