package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterProficientSkills() {
	charactersProficientSkills := []models.CharacterProficientSkill{
		{
			CharacterID:     1,
			SkillName:       "Arcana",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     1,
			SkillName:       "Athletics",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     1,
			SkillName:       "History",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     1,
			SkillName:       "Insight",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     1,
			SkillName:       "Survival",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     2,
			SkillName:       "Athletics",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     2,
			SkillName:       "Intimidation",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     2,
			SkillName:       "Perception",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     2,
			SkillName:       "Survival",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     3,
			SkillName:       "Athletics",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     3,
			SkillName:       "Insight",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     3,
			SkillName:       "Intimidation",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     3,
			SkillName:       "Investigation",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     3,
			SkillName:       "Perception",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     3,
			SkillName:       "Survival",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     4,
			SkillName:       "Insight",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     4,
			SkillName:       "Intimidation",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     4,
			SkillName:       "Perception",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     4,
			SkillName:       "Stealth",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     5,
			SkillName:       "Arcana",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     5,
			SkillName:       "Deception",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     5,
			SkillName:       "Persuasion",
			ProficiencyType: "Proficiency",
		},
		{
			CharacterID:     5,
			SkillName:       "Survival",
			ProficiencyType: "Proficiency",
		},
	}

	for _, characterProficientSkill := range charactersProficientSkills {
		err := s.DB.Where("character_id = ? AND skill_name = ?", characterProficientSkill.CharacterID, characterProficientSkill.SkillName).FirstOrCreate(&characterProficientSkill).Error
		if err != nil {
			log.Printf("error creating character proficient skill for CharacterID: %q, Skill: %s: %s", characterProficientSkill.CharacterID, characterProficientSkill.SkillName, err.Error())
		}
	}
}
