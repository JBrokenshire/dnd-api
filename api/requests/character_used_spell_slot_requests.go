package requests

type UpdateCharacterUsedSpellSlotRequest struct {
	SpellSlotsUsed int `json:"spell_slots_used" validate:"min=0"`
}
