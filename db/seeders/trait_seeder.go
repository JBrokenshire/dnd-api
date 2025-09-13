package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetTraits() {
	traits := []*m.Trait{
		{
			ID:          "earth-genasi-asi",
			Name:        "Ability Score Increase",
			Description: "Your Constitution score increases by 2 and your Strength score increases by 1.",
		},
		{
			ID:          "earth-genasi-earth-walk",
			Name:        "Earth Walk",
			Description: "You can move freely across difficult terrain made of earth or stone without expending extra movement.",
		},
		{
			ID:          "earth-genasi-merge-with-stone",
			Name:        "Merge with Stone",
			Description: "You can cast Pass Without Trace (without material components) once per long rest. CON is your spellcasting ability.",
		},
		{
			ID:          "dragonborn-asi",
			Name:        "Ability Score Increase",
			Description: "Your Strength score increases by 2 and your Charisma score increases by 1.",
		},
		{
			ID:          "dragonborn-draconic-ancestry",
			Name:        "Draconic Ancestry",
			Description: "You gain a breath weapon and damage resistance with your chosen dragon type.",
		},
		{
			ID:          "dragonborn-breath-weapon",
			Name:        "Breath Weapon",
			Description: "Once per short rest as an action, exhale destructive energy based on your Draconic Ancestry. Each creature in the area must make a DC 16 saving throw  (type determined by your ancestry), taking 2d6 ([6th] 3d6, [11th] 4d6, [16th] 5d6) on a failed save, and half damage on a successful one.",
		},
		{
			ID:          "dragonborn-damage-resistance",
			Name:        "Damage Resistance",
			Description: "You have resistance to the damage type associated with your draconic ancestry.",
		},
		{
			ID:          "stout-halfling-asi",
			Name:        "Ability Score Increase",
			Description: "Your Dexterity score increases by 2 and your Constitution score increases by 1.",
		},
		{
			ID:          "stout-halfling-lucky",
			Name:        "Lucky",
			Description: "When you roll a 1 on the D20 for an attack roll, ability check or saving throw, you can reroll the die and must use the new roll.",
		},
		{
			ID:          "stout-halfling-brave",
			Name:        "Brave",
			Description: "You have advantage on saving throws against being frightened.",
		},
		{
			ID:          "stout-halfling-halfling-nimbleness",
			Name:        "Halfling Nimbleness",
			Description: "You can move through the space of any creature that is of a size larger than yours.",
		},
		{
			ID:          "stout-halfling-stout-resilience",
			Name:        "Stout Resilience",
			Description: "You have advantage on saving throws against poison, and you have resistance to poison damage.",
		},
	}

	for _, trait := range traits {
		err := s.DB.Where("id = ?", trait.ID).FirstOrCreate(&trait).Error
		if err != nil {
			log.Printf("Error creating trait with id %v in seeder: %v", trait.ID, err.Error())
		}
	}
}
