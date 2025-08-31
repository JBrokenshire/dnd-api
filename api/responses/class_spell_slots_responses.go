package responses

import m "dnd-api/db/models"

type ClassSpellSlotsResponse struct {
	ClassLevel int `json:"class_level"`
	SpellLevel int `json:"spell_level"`
	SpellSlots int `json:"spell_slots"`
}

func NewClassSpellSlotsResponse(spellSlots *m.ClassSpellSlots) *ClassSpellSlotsResponse {
	return &ClassSpellSlotsResponse{
		ClassLevel: spellSlots.ClassLevel,
		SpellLevel: spellSlots.SpellLevel,
		SpellSlots: spellSlots.SpellSlots,
	}
}

func NewClassSpellSlotsResponses(spellSlots []*m.ClassSpellSlots) []ClassSpellSlotsResponse {
	var res []ClassSpellSlotsResponse
	for _, spellSlot := range spellSlots {
		res = append(res, *NewClassSpellSlotsResponse(spellSlot))
	}
	return res
}
