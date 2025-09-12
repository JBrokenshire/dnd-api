package responses

import m "dnd-api/db/models"

type CharacterInventoryItemResponse struct {
	ID       uint   `json:"id"`
	Location string `json:"location"`
	Quantity int    `json:"quantity"`
	Equipped *bool  `json:"equipped"`

	Item ItemResponse `json:"item"`
}

func NewCharacterInventoryItemResponse(item *m.CharacterInventoryItem) *CharacterInventoryItemResponse {
	res := &CharacterInventoryItemResponse{
		ID:       item.ID,
		Location: item.Location,
		Quantity: item.Quantity,
		Equipped: item.Equipped,
	}
	if item.Item.ID != 0 {
		res.Item = *NewItemResponse(&item.Item)
	}
	return res
}

func NewCharacterInventoryItemResponses(items []*m.CharacterInventoryItem) []CharacterInventoryItemResponse {
	var res []CharacterInventoryItemResponse
	for _, item := range items {
		res = append(res, *NewCharacterInventoryItemResponse(item))
	}
	return res
}
