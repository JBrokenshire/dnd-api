package responses

import m "dnd-api/db/models"

type ClassSpellLevelResponse struct {
	SpellLevel    int `json:"spell_level"`
	NumberOfSlots int `json:"number_of_slots"`
}

func NewClassSpellLevelResponse(classSpellLevel *m.ClassSpellLevel) *ClassSpellLevelResponse {
	return &ClassSpellLevelResponse{
		SpellLevel:    classSpellLevel.SpellLevel,
		NumberOfSlots: classSpellLevel.NumberOfSlots,
	}
}

func NewCLassSpellLevelResponses(classSpellLevels []*m.ClassSpellLevel) []ClassSpellLevelResponse {
	var res []ClassSpellLevelResponse
	for _, classSpellLevel := range classSpellLevels {
		res = append(res, *NewClassSpellLevelResponse(classSpellLevel))
	}
	return res
}
