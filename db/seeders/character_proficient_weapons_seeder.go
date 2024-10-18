package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharactersProficientWeapons() {
	charactersProficientWeapons := []models.CharacterProficientWeapon{
		{
			CharacterID: 1,
			Weapon:      "Martial Weapons",
		},
		{
			CharacterID: 1,
			Weapon:      "Simple Weapons",
		},
		{
			CharacterID: 2,
			Weapon:      "Martial Weapons",
		},
		{
			CharacterID: 2,
			Weapon:      "Simple Weapons",
		},
		{
			CharacterID: 3,
			Weapon:      "Martial Weapons",
		},
		{
			CharacterID: 3,
			Weapon:      "Simple Weapons",
		},
		{
			CharacterID: 4,
			Weapon:      "Martial Weapons",
		},
		{
			CharacterID: 4,
			Weapon:      "Simple Weapons",
		},
		{
			CharacterID: 5,
			Weapon:      "Simple Weapons",
		},
	}

	for _, characterProficientWeapon := range charactersProficientWeapons {
		err := s.DB.Where("character_id = ? AND weapon = ?", characterProficientWeapon.CharacterID, characterProficientWeapon.Weapon).FirstOrCreate(&characterProficientWeapon).Error
		if err != nil {
			log.Printf("error creating character proficient skill for CharacterID: %q, Skill: %s: %s", characterProficientWeapon.CharacterID, characterProficientWeapon.Weapon, err.Error())
		}
	}
}
