package seeders

import (
	m "dnd-api/db/models"
	"dnd-api/pkg/utils"
	"log"
)

func (s *Seeder) SetWeapons() {
	weapons := []*m.Weapon{
		{
			ItemId:        1,
			WeaponType:    m.WeaponTypeMelee,
			Distance:      20,
			AltDistance:   utils.IntPointer(60),
			Ability:       m.WeaponAbilityDEX,
			Damage:        "1d4",
			DamageType:    m.DamageTypePiercing,
			Proficiencies: `["Simple Weapons","Dagger"]`,
		},
		{
			ItemId:        2,
			WeaponType:    m.WeaponTypeMelee,
			Distance:      5,
			Ability:       m.WeaponAbilityDEX,
			Damage:        "1d8",
			DamageType:    m.DamageTypePiercing,
			Bonus:         1,
			Proficiencies: `["Martial Weapons","Rapier"]`,
		},
		{
			ItemId:        18,
			WeaponType:    m.WeaponTypeMelee,
			Distance:      5,
			Ability:       m.WeaponAbilitySTR,
			Damage:        "1d12",
			DamageType:    m.DamageTypeSlashing,
			Proficiencies: `["Martial Weapons","Greataxe"]`,
		},
		{
			ItemId:        19,
			WeaponType:    m.WeaponTypeMelee,
			Distance:      20,
			AltDistance:   utils.IntPointer(60),
			Ability:       m.WeaponAbilitySTR,
			Damage:        "1d6",
			DamageType:    m.DamageTypeSlashing,
			Proficiencies: `["Simple Weapons","Handaxe"]`,
		},
		{
			ItemId:        21,
			WeaponType:    m.WeaponTypeMelee,
			Distance:      5,
			Ability:       m.WeaponAbilitySTR,
			Damage:        "1d8",
			DamageType:    m.DamageTypePiercing,
			Proficiencies: `["Martial Weapons","Morningstar"]`,
		},
		{
			ItemId:        30,
			WeaponType:    m.WeaponTypeMelee,
			Distance:      20,
			AltDistance:   utils.IntPointer(60),
			Ability:       m.WeaponAbilityDEX,
			Damage:        "1d4",
			DamageType:    m.DamageTypePiercing,
			Bonus:         1,
			Proficiencies: `["Simple Weapons","Dagger"]`,
		},
		{
			ItemId:        33,
			WeaponType:    m.WeaponTypeRanged,
			Distance:      80,
			AltDistance:   utils.IntPointer(320),
			Ability:       m.WeaponAbilityDEX,
			Damage:        "1d8",
			DamageType:    m.DamageTypePiercing,
			Proficiencies: `["Simple Weapons","Crossbow, Light"]`,
		},
	}

	for _, weapon := range weapons {
		err := s.DB.Where("item_id = ?", weapon.ItemId).FirstOrCreate(&weapon).Error
		if err != nil {
			log.Printf("Error creating weapon with item_id %v in seeder: %v", weapon.ItemId, err.Error())
		}
	}
}
