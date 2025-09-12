package responses

import (
	m "dnd-api/db/models"
)

type ItemResponse struct {
	Name       string  `json:"name"`
	Rarity     string  `json:"rarity"`
	Notes      string  `json:"notes"`
	Cost       float32 `json:"cost"`
	Weight     float32 `json:"weight"`
	Equippable bool    `json:"equippable"`
	Origin     string  `json:"origin"`
	Type       string  `json:"type"`

	// Bonuses
	StrengthBonus         int     `json:"strength_bonus"`
	DexterityBonus        int     `json:"dexterity_bonus"`
	ConstitutionBonus     int     `json:"constitution_bonus"`
	IntelligenceBonus     int     `json:"intelligence_bonus"`
	WisdomBonus           int     `json:"wisdom_bonus"`
	CharismaBonus         int     `json:"charisma_bonus"`
	StrengthSaveBonus     int     `json:"strength_save_bonus"`
	DexteritySaveBonus    int     `json:"dexterity_save_bonus"`
	ConstitutionSaveBonus int     `json:"constitution_save_bonus"`
	IntelligenceSaveBonus int     `json:"intelligence_save_bonus"`
	WisdomSaveBonus       int     `json:"wisdom_save_bonus"`
	CharismaSaveBonus     int     `json:"charisma_save_bonus"`
	SavingThrowBonusType  *string `json:"saving_throw_bonus_type"`
	SavingThrowBonusText  string  `json:"saving_throw_bonus_text"`
	InitiativeBonus       int     `json:"initiative_bonus"`
	InitiativeAdvantage   bool    `json:"initiative_advantage"`
	ArmourClassBonus      int     `json:"armour_class_bonus"`

	Armour ArmourResponse `json:"armour"`
}

func NewItemResponse(item *m.Item) *ItemResponse {
	res := &ItemResponse{
		Name:       item.Name,
		Rarity:     item.Rarity,
		Notes:      item.Notes,
		Cost:       item.Cost,
		Weight:     item.Weight,
		Equippable: item.Equippable,
		Origin:     item.Origin,
		Type:       item.Type,

		StrengthBonus:         item.StrengthBonus,
		DexterityBonus:        item.DexterityBonus,
		ConstitutionBonus:     item.ConstitutionBonus,
		IntelligenceBonus:     item.IntelligenceBonus,
		WisdomBonus:           item.WisdomBonus,
		CharismaBonus:         item.CharismaBonus,
		StrengthSaveBonus:     item.StrengthSaveBonus,
		DexteritySaveBonus:    item.DexteritySaveBonus,
		ConstitutionSaveBonus: item.ConstitutionSaveBonus,
		IntelligenceSaveBonus: item.IntelligenceSaveBonus,
		WisdomSaveBonus:       item.WisdomSaveBonus,
		CharismaSaveBonus:     item.CharismaSaveBonus,
		SavingThrowBonusType:  item.SavingThrowBonusType,
		SavingThrowBonusText:  item.SavingThrowBonusText,
		InitiativeBonus:       item.InitiativeBonus,
		InitiativeAdvantage:   item.InitiativeAdvantage,
		ArmourClassBonus:      item.ArmourClassBonus,
	}

	if item.Type == m.ItemTypeArmour {
		res.Armour = *NewArmourResponse(&item.Armour)
	}

	return res
}
