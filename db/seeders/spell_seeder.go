package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetSpells() {
	effectUtility := "Utility"
	effectCommunication := "Communication"
	effectControl := "Control"
	effectSocial := "Social"
	effectBuff := "Buff"
	effectWarding := "Warding"
	effectShapechanging := "Shapechanging"
	effectAttack := "Attack"
	effectNegation := "Negation"
	effectFrightened := "Frightened"

	damage1d12 := "1d12"
	damage1d10 := "1d10"
	damage1d8 := "1d8"

	damageTypePoison := "Poison"
	damageTypeLightning := "Lightning"
	damageTypeThunder := "Thunder"
	damageTypePiercing := "Piercing"

	saveWisdom := "WIS"
	saveConstitution := "CON"

	spells := []models.Spell{
		{
			ID:          1,
			Name:        "Dancing Lights",
			School:      "Illusion",
			CastingTime: "1A",
			Distance:    "120",
			Effect:      &effectUtility,
			Notes:       "D: 1m, V/S/M",
		},
		{
			ID:          2,
			Name:        "Message",
			School:      "Transmutation",
			CastingTime: "1A",
			Distance:    "120",
			Effect:      &effectCommunication,
			Notes:       "D: 1Rnd, S/M",
		},
		{
			ID:          3,
			Name:        "Poison Spray",
			School:      "Necromancy",
			CastingTime: "1A",
			Distance:    "30",
			Effect:      &effectAttack,
			Damage:      &damage1d12,
			DamageType:  &damageTypePoison,
			Notes:       "V/S",
		},
		{
			ID:          4,
			Name:        "Shape Water",
			School:      "Transmutation",
			CastingTime: "1A",
			Distance:    "30",
			Effect:      &effectControl,
			Notes:       "5ft. Cube, S",
		},
		{
			ID:          5,
			Name:        "Shocking Grasp",
			School:      "Evocation",
			CastingTime: "1A",
			Distance:    "Touch",
			Effect:      &effectAttack,
			Damage:      &damage1d8,
			DamageType:  &damageTypeLightning,
			Notes:       "V/S",
		},
		{
			ID:          6,
			Name:        "Booming Blade",
			School:      "Evocation",
			CastingTime: "1A",
			Distance:    "Self",
			Damage:      &damage1d8,
			DamageType:  &damageTypeThunder,
			Notes:       "D: 1Rnd, 5ft. Sphere, S/M",
		},
		{
			ID:          7,
			Name:        "Comprehend Languages",
			School:      "Divination",
			Level:       1,
			CastingTime: "1A",
			Distance:    "Self",
			Effect:      &effectSocial,
			Notes:       "D: 1hr, V/S/M",
		},
		{
			ID:          8,
			Name:        "Feather Fall",
			School:      "Transmutation",
			Level:       1,
			CastingTime: "1R",
			Distance:    "60",
			Effect:      &effectUtility,
			Notes:       "D: 1m, V/M",
		},
		{
			ID:          9,
			Name:        "Ice Knife",
			School:      "Conjuration",
			Level:       1,
			CastingTime: "1A",
			Distance:    "60",
			Effect:      &effectAttack,
			Damage:      &damage1d10,
			DamageType:  &damageTypePiercing,
			Notes:       "5ft. Sphere, S/M",
		},
		{
			ID:          10,
			Name:        "Mage Armour",
			School:      "Abjuration",
			Level:       1,
			CastingTime: "1A",
			Distance:    "Touch",
			Effect:      &effectBuff,
			Notes:       "D: 8hr, V/S/M",
		},
		{
			ID:          11,
			Name:        "Shield",
			School:      "Abjuration",
			Level:       1,
			CastingTime: "1R",
			Distance:    "Self",
			Effect:      &effectWarding,
			Notes:       "D: 1Rnd, V/S",
		},
		{
			ID:          12,
			Name:        "Alter Self",
			School:      "Transmutation",
			Level:       2,
			CastingTime: "1A",
			Distance:    "Self",
			Effect:      &effectShapechanging,
			Notes:       "D: 1hr, V/S",
		},
		{
			ID:          13,
			Name:        "Suggestion",
			School:      "Enchantment",
			Level:       2,
			CastingTime: "1A",
			Distance:    "30",
			Effect:      &effectControl,
			Save:        &saveWisdom,
			Notes:       "D: 8hr, V/M",
		},
		{
			ID:          14,
			Name:        "Fog Cloud",
			School:      "Conjuration",
			Level:       1,
			CastingTime: "1A",
			Distance:    "120",
			Effect:      &effectControl,
			Notes:       "D: 1hr, 20ft. Sphere, V/S",
		},
		{
			ID:          15,
			Name:        "Counterspell",
			School:      "Abjuration",
			Level:       3,
			CastingTime: "1R",
			Distance:    "60",
			Effect:      &effectNegation,
			Save:        &saveConstitution,
			Notes:       "S",
		},
		{
			ID:          16,
			Name:        "Fear",
			School:      "Illusion",
			Level:       3,
			CastingTime: "1A",
			Distance:    "Self",
			Effect:      &effectFrightened,
			Save:        &saveWisdom,
			Notes:       "D: 1m, 30ft. Cone, V/S/M",
		},
	}

	for _, spell := range spells {
		err := s.DB.Where("id = ?", spell.ID).FirstOrCreate(&spell).Error
		if err != nil {
			log.Printf("error creating spell with id %v - %v", spell.ID, err)
		}
	}
}
