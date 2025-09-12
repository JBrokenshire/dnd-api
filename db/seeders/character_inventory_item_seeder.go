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
			Quantity:    1,
		},
		{
			ID:          2,
			CharacterId: 1,
			ItemId:      1,
			Location:    "Equipment",
			Equipped:    utils.BoolPointer(false),
			Quantity:    1,
		},
		{
			ID:          3,
			CharacterId: 1,
			ItemId:      2,
			Location:    "Equipment",
			Equipped:    utils.BoolPointer(true),
			Quantity:    1,
		},
		{
			ID:          4,
			CharacterId: 1,
			ItemId:      3,
			Location:    "Equipment",
			Equipped:    utils.BoolPointer(true),
			Quantity:    1,
		},
		{
			ID:          5,
			CharacterId: 1,
			ItemId:      4,
			Location:    "Equipment",
			Quantity:    1,
		},
		{
			ID:          6,
			CharacterId: 1,
			ItemId:      5,
			Location:    "Backpack",
			Quantity:    1,
		},
		{
			ID:          7,
			CharacterId: 1,
			ItemId:      6,
			Location:    "Backpack",
			Quantity:    1,
		},
		{
			ID:          8,
			CharacterId: 1,
			ItemId:      7,
			Location:    "Backpack",
			Quantity:    1,
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
			Quantity:    1,
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
			Quantity:    1,
		},
		{
			ID:          14,
			CharacterId: 2,
			ItemId:      14,
			Location:    "Equipment",
			Quantity:    1,
			Equipped:    utils.BoolPointer(true),
		},
		{
			ID:          15,
			CharacterId: 2,
			ItemId:      15,
			Location:    "Equipment",
			Quantity:    1,
			Equipped:    utils.BoolPointer(true),
		},
		{
			ID:          16,
			CharacterId: 2,
			ItemId:      16,
			Location:    "Equipment",
			Quantity:    1,
			Equipped:    utils.BoolPointer(true),
		},
		{
			ID:          17,
			CharacterId: 2,
			ItemId:      17,
			Location:    "Equipment",
			Quantity:    1,
		},
		{
			ID:          18,
			CharacterId: 2,
			ItemId:      18,
			Location:    "Equipment",
			Quantity:    1,
			Equipped:    utils.BoolPointer(true),
		},
		{
			ID:          19,
			CharacterId: 2,
			ItemId:      19,
			Location:    "Equipment",
			Quantity:    1,
			Equipped:    utils.BoolPointer(true),
		},
		{
			ID:          20,
			CharacterId: 2,
			ItemId:      19,
			Location:    "Equipment",
			Quantity:    1,
			Equipped:    utils.BoolPointer(false),
		},
		{
			ID:          21,
			CharacterId: 2,
			ItemId:      20,
			Location:    "Equipment",
			Quantity:    1,
		},
		{
			ID:          22,
			CharacterId: 2,
			ItemId:      21,
			Location:    "Equipment",
			Quantity:    1,
			Equipped:    utils.BoolPointer(true),
		},
		{
			ID:          23,
			CharacterId: 2,
			ItemId:      21,
			Location:    "Equipment",
			Quantity:    1,
			Equipped:    utils.BoolPointer(false),
		},
		{
			ID:          24,
			CharacterId: 2,
			ItemId:      22,
			Location:    "Backpack",
			Quantity:    1,
		},
		{
			ID:          25,
			CharacterId: 2,
			ItemId:      23,
			Location:    "Backpack",
			Quantity:    1,
		},
		{
			ID:          26,
			CharacterId: 2,
			ItemId:      24,
			Location:    "Backpack",
			Quantity:    1,
		},
		{
			ID:          27,
			CharacterId: 2,
			ItemId:      25,
			Location:    "Backpack",
			Quantity:    1,
		},
		{
			ID:          28,
			CharacterId: 2,
			ItemId:      26,
			Location:    "Backpack",
			Quantity:    10,
		},
		{
			ID:          29,
			CharacterId: 2,
			ItemId:      11,
			Location:    "Backpack",
			Quantity:    10,
		},
		{
			ID:          30,
			CharacterId: 2,
			ItemId:      29,
			Location:    "Backpack",
			Quantity:    1,
		},
		{
			ID:          31,
			CharacterId: 2,
			ItemId:      12,
			Location:    "Backpack",
			Quantity:    1,
		},
		{
			ID:          32,
			CharacterId: 2,
			ItemId:      27,
			Location:    "Backpack",
			Quantity:    1,
		},
		{
			ID:          33,
			CharacterId: 2,
			ItemId:      28,
			Location:    "Backpack",
			Quantity:    10,
		},
		{
			ID:          34,
			CharacterId: 2,
			ItemId:      13,
			Location:    "Backpack",
			Quantity:    1,
		},
	}

	for _, item := range characterInventoryItems {
		err := s.DB.Where("id = ?", item.ID).FirstOrCreate(&item).Error
		if err != nil {
			log.Printf("Error creating character inventory item with id %v in seeder: %v", item.ID, err.Error())
		}
	}
}
