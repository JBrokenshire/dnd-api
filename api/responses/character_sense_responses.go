package responses

import m "dnd-api/db/models"

type CharacterSenseResponse struct {
	Sense string `json:"sense"`
}

func NewCharacterSenseResponse(sense *m.CharacterSense) *CharacterSenseResponse {
	return &CharacterSenseResponse{
		Sense: sense.Sense,
	}
}

func NewCharacterSenseResponses(senses []m.CharacterSense) []CharacterSenseResponse {
	var res []CharacterSenseResponse
	for _, sense := range senses {
		res = append(res, *NewCharacterSenseResponse(&sense))
	}
	return res
}
