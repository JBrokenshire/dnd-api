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
			ShortDescription:    "Bards are expert at inspiring others, soothing hurts, disheartening foes, and creating illusions.",
			PrimaryAbility:      "Charisma",
			HitPointDieValue:    8,
			Saves:               `["Dexterity","Charisma"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityCharisma),
		},
		{
			ID:               2,
			Name:             "Fighter",
			ShortDescription: "Fighters all share an unparalleled prowess with weapons and armor, and are well acquainted with death, both meting it out and defying it.",
			PrimaryAbility:   "Strength or Dexterity",
			HitPointDieValue: 10,
			Saves:            `["Strength", "Constitution"]`,
		},
	}

	for _, class := range classes {
		err := s.DB.Where("id = ?", class.ID).FirstOrCreate(&class).Error
		if err != nil {
			log.Printf("Error creating class with id %v in seeder: %v", class.ID, err.Error())
		}
	}
}
