package tests

import (
	"dnd-api/api/requests"
	"dnd-api/db/factories"
	m "dnd-api/db/models"
	"dnd-api/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestCharacterInventoryItemUpdate(t *testing.T) {
	ts.ClearTable("characters")
	ts.ClearTable("character_inventory_items")

	// Create characters
	character := &m.Character{}
	factories.NewCharacter(ts.S.Db, character)
	character2 := &m.Character{}
	factories.NewCharacter(ts.S.Db, character2)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	// Create character inventory items
	characterInventoryItem := &m.CharacterInventoryItem{CharacterId: character.ID}
	factories.NewCharacterInventoryItem(ts.S.Db, characterInventoryItem)
	character2InventoryItem := &m.CharacterInventoryItem{CharacterId: character2.ID}
	factories.NewCharacterInventoryItem(ts.S.Db, character2InventoryItem)
	differentUserCharacterInventoryItem := &m.CharacterInventoryItem{CharacterId: differentUserCharacter.ID}
	factories.NewCharacterInventoryItem(ts.S.Db, differentUserCharacterInventoryItem)

	getRequest := func(characterId, inventoryItemId interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			Url:    fmt.Sprintf("/characters/%v/inventory-item/%v", characterId, inventoryItemId),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't update inventory item for character that doesn't exist",
			Request:     getRequest(1000, characterInventoryItem.ID),
			RequestBody: requests.UpdateCharacterInventoryItemRequest{Equipped: true},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:        "Can't update inventory item for character that belongs to a different user",
			Request:     getRequest(differentUserCharacter.ID, characterInventoryItem.ID),
			RequestBody: requests.UpdateCharacterInventoryItemRequest{Equipped: true},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:        "Can't update inventory item that doesn't exist",
			Request:     getRequest(character.ID, 1000),
			RequestBody: requests.UpdateCharacterInventoryItemRequest{Equipped: true},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Inventory item not found",
			},
		},
		{
			Name:        "Can't update inventory item that belongs to a different character",
			Request:     getRequest(character.ID, character2InventoryItem.ID),
			RequestBody: requests.UpdateCharacterInventoryItemRequest{Equipped: true},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Inventory item not found",
			},
		},
		{
			Name:        "Can update inventory item",
			Request:     getRequest(character.ID, characterInventoryItem.ID),
			RequestBody: requests.UpdateCharacterInventoryItemRequest{Equipped: true},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Character inventory item updated",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}
