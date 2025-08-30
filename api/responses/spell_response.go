package responses

import m "dnd-api/db/models"

type SpellResponse struct {
	Name        string  `json:"name"`
	School      string  `json:"school"`
	Level       int     `json:"level"`
	CastingTime string  `json:"casting_time"`
	Range       string  `json:"range"`
	IsAttack    bool    `json:"is_attack"`
	IsSave      bool    `json:"is_save"`
	SaveAbility *string `json:"save_ability"`
	Effect      string  `json:"effect"`
	Damage      *string `json:"damage"`
	DamageType  *string `json:"damage_type"`
	Notes       string  `json:"notes"`
	CanUpcast   bool    `json:"can_upcast"`
}

func NewSpellResponse(spell *m.Spell) *SpellResponse {
	return &SpellResponse{
		Name:        spell.Name,
		School:      spell.School,
		Level:       spell.Level,
		CastingTime: spell.CastingTime,
		Range:       spell.Distance,
		IsAttack:    spell.IsAttack,
		IsSave:      spell.IsSave,
		SaveAbility: spell.SaveAbility,
		Effect:      spell.Effect,
		Damage:      spell.Damage,
		DamageType:  spell.DamageType,
		Notes:       spell.Notes,
		CanUpcast:   spell.CanUpcast,
	}
}

func NewSpellResponses(spells []*m.Spell) []SpellResponse {
	var res []SpellResponse
	for _, spell := range spells {
		res = append(res, *NewSpellResponse(spell))
	}
	return res
}
