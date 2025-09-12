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
			ID:                  14,
			Name:                "Artificer",
			ShortDescription:    "Masters of invention, artificers use ingenuity and magic to unlock extraordinary capabilities in objects. They see magic as a complex system waiting to be decoded and then harnessed in their spells and inventions. You can find everything you need to play one of these inventors in the next few sections.",
			PrimaryAbility:      "Intelligence",
			HitPointDieValue:    8,
			Saves:               `["Constitution","Intelligence"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityIntelligence),
		},
	}

	for _, class := range classes {
		err := s.DB.Where("id = ?", class.ID).FirstOrCreate(&class).Error
		if err != nil {
			log.Printf("Error creating class with id %v in seeder: %v", class.ID, err.Error())
		}
	}
}
