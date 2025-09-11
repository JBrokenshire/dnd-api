package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetItems() {
	items := []m.Item{
		{
			ID:         1,
			Name:       "Dagger",
			Rarity:     m.RarityCommon,
			Origin:     "Weapon",
			Weight:     1,
			Cost:       2,
			Equippable: true,
			Notes:      "Simple, Finesse, Light, Thrown, Nick, Range (20/60)",
		},
		{
			ID:         2,
			Name:       "Rapier, +1",
			Rarity:     m.RarityUncommon,
			Origin:     "Weapon",
			Weight:     2,
			Cost:       200,
			Equippable: true,
			Notes:      "Martial, Finesse, Vex",
		},
		{
			ID:         3,
			Name:       "Leather Armour",
			Rarity:     m.RarityCommon,
			Origin:     "Light Armour",
			Weight:     10,
			Cost:       10,
			Equippable: true,
			Notes:      "AC 11",
		},
		{
			ID:     4,
			Name:   "Lute",
			Rarity: m.RarityCommon,
			Origin: "Gear",
			Weight: 2,
			Cost:   35,
			Notes:  "Instrument",
		},
		{
			ID:     5,
			Name:   "Bedroll",
			Rarity: m.RarityCommon,
			Origin: "Adventuring Gear",
			Weight: 7,
			Cost:   1,
			Notes:  "Utility",
		},
		{
			ID:     6,
			Name:   "Bell",
			Rarity: m.RarityCommon,
			Origin: "Adventuring Gear",
			Weight: 0.2,
			Cost:   1,
			Notes:  "Communication, Utility",
		},
		{
			ID:     7,
			Name:   "Bullseye Lantern",
			Rarity: m.RarityCommon,
			Origin: "Adventuring Gear",
			Weight: 2,
			Cost:   10,
			Notes:  "Utility, Exploration",
		},
		{
			ID:     8,
			Name:   "Costume",
			Rarity: m.RarityCommon,
			Origin: "Adventuring Gear",
			Weight: 4,
			Cost:   5,
			Notes:  "Social, Deception, Outerwear",
		},
		{
			ID:     9,
			Name:   "Mirror",
			Rarity: m.RarityCommon,
			Origin: "Adventuring Gear",
			Weight: 0.5,
			Cost:   5,
			Notes:  "Social, Utility",
		},
		{
			ID:     10,
			Name:   "Oil",
			Rarity: m.RarityCommon,
			Origin: "Adventuring Gear",
			Weight: 1,
			Cost:   0.1,
			Notes:  "Damage, Utility, Combat",
		},
		{
			ID:     11,
			Name:   "Rations",
			Rarity: m.RarityCommon,
			Origin: "Adventuring Gear",
			Weight: 2,
			Cost:   0.5,
			Notes:  "Social, Utility, Consumable",
		},
		{
			ID:     12,
			Name:   "Tinderbox",
			Rarity: m.RarityCommon,
			Origin: "Adventuring Gear",
			Weight: 1,
			Cost:   0.5,
			Notes:  "Utility, Exploration",
		},
		{
			ID:     13,
			Name:   "Waterskin",
			Rarity: m.RarityCommon,
			Origin: "Adventuring Gear",
			Weight: 5,
			Cost:   0.2,
			Notes:  "Container",
		},
	}

	for _, item := range items {
		err := s.DB.Where("id = ?", item.ID).FirstOrCreate(&item).Error
		if err != nil {
			log.Printf("Error creating item with id %v in seeder: %v", item.ID, err.Error())
		}
	}
}
