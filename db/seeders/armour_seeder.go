package seeders

import (
	m "dnd-api/db/models"
	"dnd-api/pkg/utils"
	"log"
)

func (s *Seeder) SetArmour() {
	armours := []m.Armour{
		{
			ItemId: 3,
			BaseAC: 11,
			Type:   m.ArmourTypeLight,
		},
		{
			ItemId:              16,
			BaseAC:              16,
			StrRequirement:      utils.IntPointer(13),
			Type:                m.ArmourTypeHeavy,
			MaxDexModifier:      utils.IntPointer(0),
			StealthDisadvantage: true,
		},
	}

	for _, armour := range armours {
		err := s.DB.Where("item_id = ?", armour.ItemId).FirstOrCreate(&armour).Error
		if err != nil {
			log.Printf("Error creating armour with item_id %v in seeder: %v", armour.ItemId, err.Error())
		}
	}
}
