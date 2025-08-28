package responses

import m "dnd-api/db/models"

type CharacterProficientSkillResponse struct {
	Skill           string `json:"skill"`
	ProficiencyType string `json:"proficiency_type"`
}

func NewCharacterProficientSkillResponse(characterProficientSkill *m.CharacterProficientSkill) *CharacterProficientSkillResponse {
	return &CharacterProficientSkillResponse{
		Skill:           characterProficientSkill.Skill,
		ProficiencyType: characterProficientSkill.ProficiencyType,
	}
}

func NewCharacterProficientSkillResponses(characterProficientSkills []*m.CharacterProficientSkill) []CharacterProficientSkillResponse {
	var res []CharacterProficientSkillResponse
	for _, profSkill := range characterProficientSkills {
		res = append(res, *NewCharacterProficientSkillResponse(profSkill))
	}
	return res
}
