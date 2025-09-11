package seeders

import (
	m "dnd-api/db/models"
	"dnd-api/pkg/utils"
	"log"
)

func (s *Seeder) SetCharacterInventoryItems() {
	characterInventoryItems := []m.CharacterInventoryItem{
		{
			ID:          1,
			CharacterId: 1,
			ItemId:      1,
			Location:    "Equipment",
			Equipped:    utils.BoolPointer(false),
		},
		{
			ID:          2,
			CharacterId: 1,
			ItemId:      1,
			Location:    "Equipment",
			Equipped:    utils.BoolPointer(false),
		},
		{
			ID:          3,
			CharacterId: 1,
			ItemId:      2,
			Location:    "Equipment",
			Equipped:    utils.BoolPointer(true),
		},
		{
			ID:          4,
			CharacterId: 1,
			ItemId:      3,
			Location:    "Equipment",
			Equipped:    utils.BoolPointer(true),
		},
		{
			ID:          5,
			CharacterId: 1,
			ItemId:      4,
			Location:    "Equipment",
		},
		{
			ID:          6,
			CharacterId: 1,
			ItemId:      5,
			Location:    "Backpack",
		},
		{
			ID:          7,
			CharacterId: 1,
			ItemId:      6,
			Location:    "Backpack",
		},
		{
			ID:          8,
			CharacterId: 1,
			ItemId:      7,
			Location:    "Backpack",
		},
		{
			ID:          9,
			CharacterId: 1,
			ItemId:      8,
			Location:    "Backpack",
			Quantity:    3,
		},
		{
			ID:          10,
			CharacterId: 1,
			ItemId:      9,
			Location:    "Backpack",
		},
		{
			ID:          11,
			CharacterId: 1,
			ItemId:      10,
			Location:    "Backpack",
			Quantity:    10,
		},
		{
			ID:          12,
			CharacterId: 1,
			ItemId:      11,
			Location:    "Backpack",
			Quantity:    10,
		},
		{
			ID:          13,
			CharacterId: 1,
			ItemId:      12,
			Location:    "Backpack",
		},
	}

	for _, item := range characterInventoryItems {
		err := s.DB.Where("id = ?", item.ID).FirstOrCreate(&item).Error
		if err != nil {
			log.Printf("Error creating character inventory item with id %v in seeder: %v", item.ID, err.Error())
		}
	}
}
