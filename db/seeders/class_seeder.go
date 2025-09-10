package seeders

import (
	m "dnd-api/db/models"
	"dnd-api/pkg/utils"
	"log"
)

func (s *Seeder) SetClasses() {
	classes := []m.Class{
		{
			ID:                  1,
			Name:                "Bard",
			ShortDescription:    "Bard",
			PrimaryAbility:      "Charisma",
			HitPointDieValue:    8,
			Saves:               `["Dexterity","Charisma"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityCharisma),
		},
	}

	for _, class := range classes {
		err := s.DB.Where("id = ?", class.ID).FirstOrCreate(&class).Error
		if err != nil {
			log.Printf("Error creating class with id %v in seeder: %v", class.ID, err.Error())
		}
	}
}
