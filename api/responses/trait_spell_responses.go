package responses

import m "dnd-api/db/models"

type TraitSpellResponse struct {
	Uses  *int    `json:"uses"`
	Reset *string `json:"reset"`

	Spell SpellResponse `json:"spell"`
}

func NewTraitSpellResponse(traitSpell *m.TraitSpell) *TraitSpellResponse {
	res := &TraitSpellResponse{
		Uses:  traitSpell.Uses,
		Reset: traitSpell.Reset,
	}
	if traitSpell.Spell.ID != 0 {
		res.Spell = *NewSpellResponse(&traitSpell.Spell)
	}

	return res
}

func NewTraitSpellResponses(traitSpells []*m.TraitSpell) []TraitSpellResponse {
	var res []TraitSpellResponse
	for _, traitSpell := range traitSpells {
		res = append(res, *NewTraitSpellResponse(traitSpell))
	}
	return res
}
