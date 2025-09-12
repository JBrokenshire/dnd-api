package tests

import (
	"dnd-api/api/requests"
	"dnd-api/db/factories"
	m "dnd-api/db/models"
	"dnd-api/pkg/utils"
	"dnd-api/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestCharacterInventoryItemUpdate(t *testing.T) {
	ts.ClearTable("characters")
	ts.ClearTable("character_inventory_items")
	ts.ClearTable("items")
	ts.SetupDefaultUsers()

	// Create characters
	character := &m.Character{}
	factories.NewCharacter(ts.S.Db, character)
	character2 := &m.Character{}
	factories.NewCharacter(ts.S.Db, character2)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	// Create items
	item := &m.Item{}
	factories.NewItem(ts.S.Db, item)
	armourItem := &m.Item{Type: m.ItemTypeArmour}
	factories.NewItem(ts.S.Db, armourItem)
	armourItem2 := &m.Item{Type: m.ItemTypeArmour}
	factories.NewItem(ts.S.Db, armourItem2)

	// Create character inventory items
	characterInventoryItem := &m.CharacterInventoryItem{CharacterId: character.ID, ItemId: item.ID}
	factories.NewCharacterInventoryItem(ts.S.Db, characterInventoryItem)
	character2InventoryItem := &m.CharacterInventoryItem{CharacterId: character2.ID, ItemId: item.ID}
	factories.NewCharacterInventoryItem(ts.S.Db, character2InventoryItem)
	differentUserCharacterInventoryItem := &m.CharacterInventoryItem{CharacterId: differentUserCharacter.ID, ItemId: item.ID}
	factories.NewCharacterInventoryItem(ts.S.Db, differentUserCharacterInventoryItem)
	characterEquippedArmourInventoryItem := &m.CharacterInventoryItem{CharacterId: character.ID, ItemId: armourItem.ID, Equipped: utils.BoolPointer(true)}
	factories.NewCharacterInventoryItem(ts.S.Db, characterEquippedArmourInventoryItem)
	characterUnequippedArmourInventoryItem := &m.CharacterInventoryItem{CharacterId: character.ID, ItemId: armourItem2.ID, Equipped: utils.BoolPointer(false)}
	factories.NewCharacterInventoryItem(ts.S.Db, characterUnequippedArmourInventoryItem)

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
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Inventory item was updated",
					Model: m.CharacterInventoryItem{
						ID:       characterInventoryItem.ID,
						Equipped: utils.BoolPointer(true),
					},
					CountExpected: 1,
				},
			},
		},
		{
			Name:        "Can unequip armour when equipping new armour",
			Request:     getRequest(character.ID, characterUnequippedArmourInventoryItem.ID),
			RequestBody: requests.UpdateCharacterInventoryItemRequest{Equipped: true},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Character inventory item updated",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Name: "New armour was equipped",
						Model: m.CharacterInventoryItem{
							ID:       characterUnequippedArmourInventoryItem.ID,
							Equipped: utils.BoolPointer(true),
						},
						CountExpected: 1,
					},
					{
						Name: "Old armour was unequipped",
						Model: m.CharacterInventoryItem{
							ID:       characterEquippedArmourInventoryItem.ID,
							Equipped: utils.BoolPointer(false),
						},
						CountExpected: 1,
					},
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}
