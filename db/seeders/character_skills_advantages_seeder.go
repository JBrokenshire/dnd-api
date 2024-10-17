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
			SkillName:   "Arcana",
			Advantage:   true,
		},
		{
			ID:          2,
			CharacterID: 1,
			SkillName:   "History",
			Advantage:   true,
		},
		{
			ID:          3,
			CharacterID: 1,
			SkillName:   "Investigation",
			Advantage:   true,
		},
		{
			ID:          4,
			CharacterID: 1,
			SkillName:   "Nature",
			Advantage:   true,
		},
		{
			ID:          5,
			CharacterID: 1,
			SkillName:   "Religion",
			Advantage:   true,
		},
		{
			ID:          6,
			CharacterID: 1,
			SkillName:   "Survival",
			Advantage:   true,
		},
		{
			ID:           7,
			CharacterID:  2,
			SkillName:    "Stealth",
			Disadvantage: true,
		},
		{
			ID:          8,
			CharacterID: 3,
			SkillName:   "Arcana",
			Advantage:   true,
		},
		{
			ID:          9,
			CharacterID: 3,
			SkillName:   "History",
			Advantage:   true,
		},
		{
			ID:          10,
			CharacterID: 3,
			SkillName:   "Investigation",
			Advantage:   true,
		},
		{
			ID:          11,
			CharacterID: 3,
			SkillName:   "Nature",
			Advantage:   true,
		},
		{
			ID:          12,
			CharacterID: 3,
			SkillName:   "Perception",
			Advantage:   true,
		},
		{
			ID:          13,
			CharacterID: 3,
			SkillName:   "Religion",
			Advantage:   true,
		},
		{
			ID:           14,
			CharacterID:  3,
			SkillName:    "Stealth",
			Disadvantage: true,
		},
		{
			ID:          15,
			CharacterID: 3,
			SkillName:   "Survival",
			Advantage:   true,
		},
		{
			ID:           16,
			CharacterID:  4,
			SkillName:    "Stealth",
			Disadvantage: true,
		},
	}

	for _, advantage := range characterSkillsAdvantages {
		err := s.DB.Table("character_skills_advantages").Where("id = ?", advantage.ID).FirstOrCreate(&advantage).Error
		if err != nil {
			log.Printf("error creating character skills advantage with CharacterID: %v, SkillName: %v", advantage.CharacterID, advantage.SkillName)
		}
	}
}
