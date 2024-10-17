package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterSkillsAdvantages() {
	characterSkillsAdvantages := []models.CharacterSkillAdvantage{
		{
			ID:          1,
			CharacterID: 1,
			SkillName:   "Acrobatics",
		},
		{
			ID:          2,
			CharacterID: 1,
			SkillName:   "Animal Handling",
		},
		{
			ID:          3,
			CharacterID: 1,
			SkillName:   "Arcana",
			Advantage:   true,
		},
		{
			ID:          4,
			CharacterID: 1,
			SkillName:   "Athletics",
		},
		{
			ID:          5,
			CharacterID: 1,
			SkillName:   "Deception",
		},
		{
			ID:          6,
			CharacterID: 1,
			SkillName:   "History",
			Advantage:   true,
		},
		{
			ID:          7,
			CharacterID: 1,
			SkillName:   "Insight",
		},
		{
			ID:          8,
			CharacterID: 1,
			SkillName:   "Intimidation",
		},
		{
			ID:          9,
			CharacterID: 1,
			SkillName:   "Investigation",
			Advantage:   true,
		},
		{
			ID:          10,
			CharacterID: 1,
			SkillName:   "Medicine",
		},
		{
			ID:          11,
			CharacterID: 1,
			SkillName:   "Nature",
			Advantage:   true,
		},
		{
			ID:          12,
			CharacterID: 1,
			SkillName:   "Perception",
		},
		{
			ID:          13,
			CharacterID: 1,
			SkillName:   "Performance",
		},
		{
			ID:          14,
			CharacterID: 1,
			SkillName:   "Persuasion",
		},
		{
			ID:          15,
			CharacterID: 1,
			SkillName:   "Religion",
			Advantage:   true,
		},
		{
			ID:          16,
			CharacterID: 1,
			SkillName:   "Sleight of Hand",
		},
		{
			ID:          17,
			CharacterID: 1,
			SkillName:   "Stealth",
		},
		{
			ID:          18,
			CharacterID: 1,
			SkillName:   "Survival",
			Advantage:   true,
		},
		{
			ID:          19,
			CharacterID: 2,
			SkillName:   "Acrobatics",
		},
		{
			ID:          20,
			CharacterID: 2,
			SkillName:   "Animal Handling",
		},
		{
			ID:          21,
			CharacterID: 2,
			SkillName:   "Arcana",
		},
		{
			ID:          22,
			CharacterID: 2,
			SkillName:   "Athletics",
		},
		{
			ID:          23,
			CharacterID: 2,
			SkillName:   "Deception",
		},
		{
			ID:          24,
			CharacterID: 2,
			SkillName:   "History",
		},
		{
			ID:          25,
			CharacterID: 2,
			SkillName:   "Insight",
		},
		{
			ID:          26,
			CharacterID: 2,
			SkillName:   "Intimidation",
		},
		{
			ID:          27,
			CharacterID: 2,
			SkillName:   "Investigation",
		},
		{
			ID:          28,
			CharacterID: 2,
			SkillName:   "Medicine",
		},
		{
			ID:          29,
			CharacterID: 2,
			SkillName:   "Nature",
		},
		{
			ID:          30,
			CharacterID: 2,
			SkillName:   "Perception",
		},
		{
			ID:          31,
			CharacterID: 2,
			SkillName:   "Performance",
		},
		{
			ID:          32,
			CharacterID: 2,
			SkillName:   "Persuasion",
		},
		{
			ID:          33,
			CharacterID: 2,
			SkillName:   "Religion",
		},
		{
			ID:          34,
			CharacterID: 2,
			SkillName:   "Sleight of Hand",
		},
		{
			ID:           35,
			CharacterID:  2,
			SkillName:    "Stealth",
			Disadvantage: true,
		},
		{
			ID:          36,
			CharacterID: 2,
			SkillName:   "Survival",
		},
		{
			ID:          37,
			CharacterID: 3,
			SkillName:   "Acrobatics",
		},
		{
			ID:          38,
			CharacterID: 3,
			SkillName:   "Animal Handling",
		},
		{
			ID:          39,
			CharacterID: 3,
			SkillName:   "Arcana",
			Advantage:   true,
		},
		{
			ID:          40,
			CharacterID: 3,
			SkillName:   "Athletics",
		},
		{
			ID:          41,
			CharacterID: 3,
			SkillName:   "Deception",
		},
		{
			ID:          42,
			CharacterID: 1,
			SkillName:   "History",
			Advantage:   true,
		},
		{
			ID:          43,
			CharacterID: 3,
			SkillName:   "Insight",
		},
		{
			ID:          44,
			CharacterID: 3,
			SkillName:   "Intimidation",
		},
		{
			ID:          45,
			CharacterID: 3,
			SkillName:   "Investigation",
			Advantage:   true,
		},
		{
			ID:          46,
			CharacterID: 3,
			SkillName:   "Medicine",
		},
		{
			ID:          47,
			CharacterID: 3,
			SkillName:   "Nature",
			Advantage:   true,
		},
		{
			ID:          48,
			CharacterID: 3,
			SkillName:   "Perception",
			Advantage:   true,
		},
		{
			ID:          49,
			CharacterID: 3,
			SkillName:   "Performance",
		},
		{
			ID:          50,
			CharacterID: 3,
			SkillName:   "Persuasion",
		},
		{
			ID:          51,
			CharacterID: 3,
			SkillName:   "Religion",
			Advantage:   true,
		},
		{
			ID:          52,
			CharacterID: 3,
			SkillName:   "Sleight of Hand",
		},
		{
			ID:           53,
			CharacterID:  3,
			SkillName:    "Stealth",
			Disadvantage: true,
		},
		{
			ID:          54,
			CharacterID: 3,
			SkillName:   "Survival",
			Advantage:   true,
		},
		{
			ID:          55,
			CharacterID: 4,
			SkillName:   "Acrobatics",
		},
		{
			ID:          56,
			CharacterID: 4,
			SkillName:   "Animal Handling",
		},
		{
			ID:          57,
			CharacterID: 4,
			SkillName:   "Arcana",
		},
		{
			ID:          58,
			CharacterID: 4,
			SkillName:   "Athletics",
		},
		{
			ID:          59,
			CharacterID: 4,
			SkillName:   "Deception",
		},
		{
			ID:          60,
			CharacterID: 4,
			SkillName:   "History",
		},
		{
			ID:          61,
			CharacterID: 4,
			SkillName:   "Insight",
		},
		{
			ID:          62,
			CharacterID: 4,
			SkillName:   "Intimidation",
		},
		{
			ID:          63,
			CharacterID: 4,
			SkillName:   "Investigation",
		},
		{
			ID:          64,
			CharacterID: 4,
			SkillName:   "Medicine",
		},
		{
			ID:          65,
			CharacterID: 4,
			SkillName:   "Nature",
		},
		{
			ID:          66,
			CharacterID: 4,
			SkillName:   "Perception",
		},
		{
			ID:          67,
			CharacterID: 4,
			SkillName:   "Performance",
		},
		{
			ID:          68,
			CharacterID: 4,
			SkillName:   "Persuasion",
		},
		{
			ID:          69,
			CharacterID: 4,
			SkillName:   "Religion",
		},
		{
			ID:          70,
			CharacterID: 4,
			SkillName:   "Sleight of Hand",
		},
		{
			ID:           71,
			CharacterID:  4,
			SkillName:    "Stealth",
			Disadvantage: true,
		},
		{
			ID:          72,
			CharacterID: 4,
			SkillName:   "Survival",
		},
	}

	for _, advantage := range characterSkillsAdvantages {
		err := s.DB.Table("character_skills_advantages").Where("id = ?", advantage.ID).FirstOrCreate(&advantage).Error
		if err != nil {
			log.Printf("error creating character skills advantage with CharacterID: %v, SkillName: %v", advantage.CharacterID, advantage.SkillName)
		}
	}
}
