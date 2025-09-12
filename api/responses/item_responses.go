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
}

func NewItemResponse(item *m.Item) *ItemResponse {
	return &ItemResponse{
		Name:       item.Name,
		Rarity:     item.Rarity,
		Notes:      item.Notes,
		Cost:       item.Cost,
		Weight:     item.Weight,
		Equippable: item.Equippable,
		Origin:     item.Origin,
	}
}

func NewItemResponses(items []*m.Item) []ItemResponse {
	var res []ItemResponse
	for _, item := range items {
		res = append(res, *NewItemResponse(item))
	}
	return res
}
