package responses

import m "dnd-api/db/models"

type CharacterUsedSpellSlotResponse struct {
	SpellLevel     int `json:"spell_level"`
	SpellSlotsUsed int `json:"spell_slots_used"`
}

func NewCharacterUsedSpellSlotResponse(usedSpellSlot *m.CharacterUsedSpellSlot) *CharacterUsedSpellSlotResponse {
	return &CharacterUsedSpellSlotResponse{
		SpellLevel:     usedSpellSlot.SpellLevel,
		SpellSlotsUsed: usedSpellSlot.SpellSlotsUsed,
	}
}

func NewCharacterUsedSpellSlotResponses(usedSpellSlots []*m.CharacterUsedSpellSlot) []CharacterUsedSpellSlotResponse {
	var res []CharacterUsedSpellSlotResponse
	for _, usedSpellSlot := range usedSpellSlots {
		res = append(res, *NewCharacterUsedSpellSlotResponse(usedSpellSlot))
	}
	return res
}
