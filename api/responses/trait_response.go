package responses

import m "dnd-api/db/models"

type TraitResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	TraitSpells []TraitSpellResponse  `json:"trait_spells"`
	Options     []TraitOptionResponse `json:"options"`
}

func NewTraitResponse(trait *m.Trait) *TraitResponse {
	res := &TraitResponse{
		Name:        trait.Name,
		Description: trait.Description,
	}
	if len(trait.TraitSpells) > 0 {
		res.TraitSpells = NewTraitSpellResponses(trait.TraitSpells)
	}
	if len(trait.Options) > 0 {
		res.Options = NewTraitOptionResponses(trait.Options)
	}

	return res
}

func NewTraitResponses(traits []*m.Trait) []TraitResponse {
	var res []TraitResponse
	for _, trait := range traits {
		res = append(res, *NewTraitResponse(trait))
	}
	return res
}
