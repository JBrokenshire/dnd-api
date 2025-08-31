package seeders

import (
	m "dnd-api/db/models"
	"dnd-api/pkg/utils"
	"log"
)

func (s *Seeder) SetClasses() {
	classes := []*m.Class{
		{
			ID:                  1,
			Name:                "Barbarian",
			ShortDescription:    "Barbarians are mighty warriors who are powered by primal forces of the multiverse that manifest as a Rage.",
			PrimaryAbility:      "Strength",
			HitPointDieValue:    12,
			Saves:               `["Strength","Constitution"]`,
			SpellcastingAbility: nil,
		},
		{
			ID:                  2,
			Name:                "Bard",
			ShortDescription:    "Bards are expert at inspiring others, soothing hurts, disheartening foes, and creating illusions.",
			PrimaryAbility:      "Charisma",
			HitPointDieValue:    8,
			Saves:               `["Charisma","Dexterity"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityCharisma),
		},
		{
			ID:                  3,
			Name:                "Cleric",
			ShortDescription:    "Clerics can reach out to the divine magic of the Outer Planes and channel it to bolster people and battle foes.",
			PrimaryAbility:      "Wisdom",
			HitPointDieValue:    8,
			Saves:               `["Charisma","Wisdom"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityWisdom),
		},
		{
			ID:                  4,
			Name:                "Druid",
			ShortDescription:    "Druids call on the forces of nature, harnessing magic to heal, transform into animals, and wield elemental destruction.",
			PrimaryAbility:      "Wisdom",
			HitPointDieValue:    8,
			Saves:               `["Intelligence","Wisdom"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityWisdom),
		},
		{
			ID:               5,
			Name:             "Fighter",
			ShortDescription: "Fighters all share an unparalleled prowess with weapons and armor, and are well acquainted with death, both meting it out and defying it.",
			PrimaryAbility:   "Strength or Dexterity",
			HitPointDieValue: 10,
			Saves:            `["Strength","Constitution"]`,
		},
		{
			ID:               6,
			Name:             "Monk",
			ShortDescription: "Monks focus their internal reservoirs of power to create extraordinary, even supernatural, effects.",
			PrimaryAbility:   "Dexterity & Wisdom",
			HitPointDieValue: 8,
			Saves:            `["Strength","Dexterity"]`,
		},
		{
			ID:                  7,
			Name:                "Paladin",
			ShortDescription:    "Paladins live on the front lines of the cosmic struggle, united by their oaths against the forces of annihilation.",
			PrimaryAbility:      "Strength & Charisma",
			HitPointDieValue:    10,
			Saves:               `["Wisdom","Charisma"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityCharisma),
		},
		{
			ID:                  8,
			Name:                "Ranger",
			ShortDescription:    "Rangers are honed with deadly focus and harness primal powers to protect the world from the ravages of monsters and tyrants.",
			PrimaryAbility:      "Dexterity & Wisdom",
			HitPointDieValue:    10,
			Saves:               `["Strength","Dexterity"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityWisdom),
		},
		{
			ID:               9,
			Name:             "Rogue",
			ShortDescription: "Rogues have a knack for finding the solution to just about any problem, prioritizing subtle strikes over brute strength.",
			PrimaryAbility:   "Dexterity",
			HitPointDieValue: 8,
			Saves:            `["Dexterity","Intelligence"]`,
		},
		{
			ID:                  10,
			Name:                "Sorcerer",
			ShortDescription:    "Sorcerers harness and channel the raw, roiling power of innate magic that is stamped into their very being.",
			PrimaryAbility:      "Charisma",
			HitPointDieValue:    6,
			Saves:               `["Constitution","Charisma"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityCharisma),
		},
		{
			ID:                  11,
			Name:                "Warlock",
			ShortDescription:    "Warlocks quest for knowledge that lies hidden in the fabric of the multiverse, piecing together arcane secrets to bolster their own power.",
			PrimaryAbility:      "Charisma",
			HitPointDieValue:    8,
			Saves:               `["Wisdom","Charisma"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityCharisma),
		},
		{
			ID:                  12,
			Name:                "Wizard",
			ShortDescription:    "Wizards cast spells of explosive fire, arcing lightning, subtle deception, and spectacular transformations.",
			PrimaryAbility:      "Intelligence",
			HitPointDieValue:    6,
			Saves:               `["Intelligence","Wisdom"]`,
			SpellcastingAbility: utils.StrPointer(m.SpellcastingAbilityIntelligence),
		},
		{
			ID:               13,
			Name:             "Blood Hunter",
			ShortDescription: "Willing to suffer whatever it takes to achieve victory, these adept warriors have forged themselves into a potent force dedicated to protecting the innocent.",
			PrimaryAbility:   "Strength or Dexterity, & Intelligence or Wisdom",
			HitPointDieValue: 10,
			Saves:            `["Dexterity","Intelligence"]`,
		},
	}

	for _, class := range classes {
		err := s.DB.Where("id = ?", class.ID).FirstOrCreate(&class).Error
		if err != nil {
			log.Printf("Error creating class with id %v in seeder: %v", class.ID, err.Error())
		}
	}
}
