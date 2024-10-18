package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharactersMoney() {
	characterMoney := []models.CharacterMoney{
		{
			ID:          1,
			CharacterID: 1,
			Money:       "gold",
			Amount:      10,
		},
		{
			ID:          2,
			CharacterID: 1,
			Money:       "silver",
			Amount:      4,
		},
		{
			ID:          3,
			CharacterID: 1,
			Money:       "copper",
			Amount:      17,
		},
		{
			ID:          4,
			CharacterID: 2,
			Money:       "gold",
			Amount:      65,
		},
		{
			ID:          5,
			CharacterID: 2,
			Money:       "silver",
			Amount:      6,
		},
		{
			ID:          6,
			CharacterID: 2,
			Money:       "copper",
			Amount:      4,
		},
		{
			ID:          7,
			CharacterID: 3,
			Money:       "gold",
			Amount:      127,
		},
		{
			ID:          8,
			CharacterID: 3,
			Money:       "electrum",
			Amount:      5,
		},
		{
			ID:          9,
			CharacterID: 3,
			Money:       "silver",
			Amount:      37,
		},
		{
			ID:          10,
			CharacterID: 3,
			Money:       "copper",
			Amount:      116,
		},
		{
			ID:          11,
			CharacterID: 4,
			Money:       "gold",
			Amount:      27,
		},
		{
			ID:          12,
			CharacterID: 4,
			Money:       "silver",
			Amount:      9,
		},
		{
			ID:          13,
			CharacterID: 4,
			Money:       "copper",
			Amount:      2,
		},
		{
			ID:          14,
			CharacterID: 5,
			Money:       "gold",
			Amount:      30,
		},
	}

	for _, money := range characterMoney {
		err := s.DB.Table("character_money").Where("id = ?", money.ID).FirstOrCreate(&money).Error
		if err != nil {
			log.Printf("error creating character money for CharacterID: %q, Money: %s: %s", money.CharacterID, money.Money, err.Error())
		}
	}
}
