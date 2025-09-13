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
		{
			ID:               3,
			Name:             "Blood Hunter",
			ShortDescription: "Willing to suffer whatever it takes to achieve victory, these adept warriors have forged themselves into a potent force dedicated to protecting the innocent.",
			PrimaryAbility:   "Strength or Dexterity, Intelligence or Wisdom",
			HitPointDieValue: 10,
			Saves:            `["Dexterity","Intelligence"]`,
		},
	}

	for _, class := range classes {
		err := s.DB.Where("id = ?", class.ID).FirstOrCreate(&class).Error
		if err != nil {
			log.Printf("Error creating class with id %v in seeder: %v", class.ID, err.Error())
		}
	}
}
