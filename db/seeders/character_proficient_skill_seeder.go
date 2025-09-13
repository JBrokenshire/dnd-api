package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterProficientSkills() {
	characterProficientSkills := []m.CharacterProficientSkill{
		{
			ID:              1,
			CharacterID:     1,
			Skill:           m.SkillAcrobatics,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              2,
			CharacterID:     1,
			Skill:           m.SkillAnimalHandling,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              3,
			CharacterID:     1,
			Skill:           m.SkillArcana,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              4,
			CharacterID:     1,
			Skill:           m.SkillAthletics,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              5,
			CharacterID:     1,
			Skill:           m.SkillDeception,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              6,
			CharacterID:     1,
			Skill:           m.SkillHistory,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              7,
			CharacterID:     1,
			Skill:           m.SkillInsight,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              8,
			CharacterID:     1,
			Skill:           m.SkillIntimidation,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              9,
			CharacterID:     1,
			Skill:           m.SkillInvestigation,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              10,
			CharacterID:     1,
			Skill:           m.SkillMedicine,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              11,
			CharacterID:     1,
			Skill:           m.SkillNature,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              12,
			CharacterID:     1,
			Skill:           m.SkillPerception,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              13,
			CharacterID:     1,
			Skill:           m.SkillPerformance,
			ProficiencyType: m.ProficiencyTypeExpertise,
		},
		{
			ID:              14,
			CharacterID:     1,
			Skill:           m.SkillPersuasion,
			ProficiencyType: m.ProficiencyTypeExpertise,
		},
		{
			ID:              15,
			CharacterID:     1,
			Skill:           m.SkillSleightOfHand,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              16,
			CharacterID:     1,
			Skill:           m.SkillStealth,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              17,
			CharacterID:     1,
			Skill:           m.SkillSurvival,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
		{
			ID:              18,
			CharacterID:     1,
			Skill:           m.SkillReligion,
			ProficiencyType: m.ProficiencyTypeHalf,
		},
		{
			ID:              19,
			CharacterID:     2,
			Skill:           m.SkillInsight,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
		{
			ID:              20,
			CharacterID:     2,
			Skill:           m.SkillIntimidation,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
		{
			ID:              21,
			CharacterID:     2,
			Skill:           m.SkillPerception,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
		{
			ID:              22,
			CharacterID:     2,
			Skill:           m.SkillStealth,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
		{
			ID:              23,
			CharacterID:     3,
			Skill:           m.SkillArcana,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
		{
			ID:              24,
			CharacterID:     3,
			Skill:           m.SkillAthletics,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
		{
			ID:              25,
			CharacterID:     3,
			Skill:           m.SkillHistory,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
		{
			ID:              26,
			CharacterID:     3,
			Skill:           m.SkillInsight,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
		{
			ID:              27,
			CharacterID:     3,
			Skill:           m.SkillSurvival,
			ProficiencyType: m.ProficiencyTypeProficiency,
		},
	}

	for _, characterProficientSkill := range characterProficientSkills {
		err := s.DB.Where("id = ?", characterProficientSkill.ID).FirstOrCreate(&characterProficientSkill).Error
		if err != nil {
			log.Printf("Error creating character proficient skill with id %v in seeder: %v", characterProficientSkill.ID, err.Error())
		}
	}
}
