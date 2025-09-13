package responses

import m "dnd-api/db/models"

type TraitOptionResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewTraitOptionResponse(traitOption *m.TraitOption) *TraitOptionResponse {
	return &TraitOptionResponse{
		Name:        traitOption.Name,
		Description: traitOption.Description,
	}
}

func NewTraitOptionResponses(traitOptions []*m.TraitOption) []TraitOptionResponse {
	var res []TraitOptionResponse
	for _, traitOption := range traitOptions {
		res = append(res, *NewTraitOptionResponse(traitOption))
	}
	return res
}
