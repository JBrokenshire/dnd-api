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

	Armour ArmourResponse `json:"armour;omitempty"`
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
	}

	if item.Type == m.ItemTypeArmour {
		res.Armour = *NewArmourResponse(&item.Armour)
	}

	return res
}
