package responses

import m "dnd-api/db/models"

type ArmourResponse struct {
	BaseAC              int    `json:"base_ac"`
	StrRequirement      *int   `json:"str_requirement"`
	Type                string `json:"type"`
	MaxDexModifier      *int   `json:"max_dex_modifier"`
	StealthDisadvantage bool   `json:"stealth_disadvantage"`
}

func NewArmourResponse(armour *m.Armour) *ArmourResponse {
	return &ArmourResponse{
		BaseAC:              armour.BaseAC,
		StrRequirement:      armour.StrRequirement,
		Type:                armour.Type,
		MaxDexModifier:      armour.MaxDexModifier,
		StealthDisadvantage: armour.StealthDisadvantage,
	}
}
