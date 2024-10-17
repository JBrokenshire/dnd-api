package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetArmour() {
	armours := []models.Armour{
		{
			ItemID:               4,
			Type:                 "Light Armour",
			BaseAC:               12,
			MaxDexterityModifier: 10,
		},
		{
			ItemID:               22,
			Type:                 "Medium Armour",
			BaseAC:               14,
			MaxDexterityModifier: 2,
			StealthDisadvantage:  true,
		},
		{
			ItemID:               25,
			Type:                 "Medium Armour",
			BaseAC:               15,
			MaxDexterityModifier: 2,
			StealthDisadvantage:  true,
		},
		{
			ItemID:               30,
			Type:                 "Heavy Armour",
			BaseAC:               16,
			StealthDisadvantage:  true,
			StrengthRequirement:  13,
			MaxDexterityModifier: 0,
		},
	}

	for _, armour := range armours {
		err := s.DB.Table("armour").Where("item_id = ?", armour.ItemID).FirstOrCreate(&armour).Error
		if err != nil {
			log.Printf("error creating armour with item id: %q -- %s", armour.ItemID, err.Error())
		}
	}
}
