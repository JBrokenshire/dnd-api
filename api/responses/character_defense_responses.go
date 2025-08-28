package responses

import m "dnd-api/db/models"

type CharacterDefenseResponse struct {
	DamageType  string `json:"damage_type"`
	DefenseType string `json:"defense_type"`
}

func NewCharacterDefenseResponse(defense *m.CharacterDefense) *CharacterDefenseResponse {
	return &CharacterDefenseResponse{
		DamageType:  defense.DamageType,
		DefenseType: defense.DefenseType,
	}
}

func NewCharacterDefenseResponses(defenses []*m.CharacterDefense) []CharacterDefenseResponse {
	var res []CharacterDefenseResponse
	for _, defense := range defenses {
		res = append(res, *NewCharacterDefenseResponse(defense))
	}
	return res
}
