package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCharacterInventory(t *testing.T) {
	ts.ClearTable("character_inventory_items")
	ts.ClearTable("characters")
	ts.ClearTable("items")

	characterOne := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, characterOne)

	characterTwo := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, characterTwo)

	noItemsCharacter := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, noItemsCharacter)

	itemOne := &models.Item{
		ID:   1,
		Name: "Test Item One",
	}
	factories.NewItem(ts.S.Db, itemOne)

	itemTwo := &models.Item{
		ID:   2,
		Name: "Test Item Two",
	}
	factories.NewItem(ts.S.Db, itemTwo)

	characterInventoryItemOne := &models.CharacterInventoryItem{
		CharacterID: characterOne.ID,
		ItemID:      itemOne.ID,
	}
	factories.NewCharacterInventoryItem(ts.S.Db, characterInventoryItemOne)

	characterInventoryItemTwo := &models.CharacterInventoryItem{
		CharacterID: characterTwo.ID,
		ItemID:      itemTwo.ID,
	}
	factories.NewCharacterInventoryItem(ts.S.Db, characterInventoryItemTwo)

	cases := []helpers.TestCase{
		{
			TestName: "Can get all inventory items for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/inventory",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%s"`, itemOne.Name),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%s"`, itemTwo.Name),
				},
			},
		},
		{
			TestName: "Empty list for character with no inventory items",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/3/inventory",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "[]",
			},
		},
		{
			TestName: "404 response for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/inventory",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestGetCharacterEquippedWeapons(t *testing.T) {
	ts.ClearTable("character_inventory_items")
	ts.ClearTable("characters")
	ts.ClearTable("weapons")
	ts.ClearTable("items")

	characterOne := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, characterOne)

	characterTwo := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, characterTwo)

	characterThree := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, characterThree)

	itemOne := &models.Item{ID: 1, Name: "Test Item One"}
	factories.NewItem(ts.S.Db, itemOne)
	weaponOne := &models.Weapon{ItemID: 1}
	factories.NewWeapon(ts.S.Db, weaponOne)

	itemTwo := &models.Item{ID: 2, Name: "Test Item Two"}
	factories.NewItem(ts.S.Db, itemTwo)
	weaponTwo := &models.Weapon{ItemID: 2}
	factories.NewWeapon(ts.S.Db, weaponTwo)

	itemThree := &models.Item{ID: 3, Name: "Test Item Three"}
	factories.NewItem(ts.S.Db, itemThree)
	weaponThree := &models.Weapon{ItemID: 3}
	factories.NewWeapon(ts.S.Db, weaponThree)

	itemFour := &models.Item{ID: 4, Name: "Test Item Four"}
	factories.NewItem(ts.S.Db, itemFour)

	equippedWeapon := &models.CharacterInventoryItem{CharacterID: characterOne.ID, ItemID: itemOne.ID, Equipped: true, Type: "weapon"}
	factories.NewCharacterInventoryItem(ts.S.Db, equippedWeapon)

	unequippedWeapon := &models.CharacterInventoryItem{CharacterID: characterTwo.ID, ItemID: itemTwo.ID, Equipped: false, Type: "weapon"}
	factories.NewCharacterInventoryItem(ts.S.Db, unequippedWeapon)

	notWeapon := &models.CharacterInventoryItem{CharacterID: characterOne.ID, ItemID: itemFour.ID}
	factories.NewCharacterInventoryItem(ts.S.Db, notWeapon)

	characterTwoEquippedWeapon := &models.CharacterInventoryItem{CharacterID: characterTwo.ID, ItemID: itemThree.ID, Equipped: true, Type: "weapon"}
	factories.NewCharacterInventoryItem(ts.S.Db, characterTwoEquippedWeapon)

	cases := []helpers.TestCase{
		{
			TestName: "Can get equipped weapons for character - not unequipped, non-weapons or other characters",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/inventory/equipped-weapons",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%s"`, itemOne.Name),
					fmt.Sprintf(`"item_id":%v`, itemOne.ID),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%s"`, itemTwo.Name),
					fmt.Sprintf(`"name":"%s"`, itemThree.Name),
					fmt.Sprintf(`"name":"%s"`, itemFour.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestGetCharacterEquippedArmour(t *testing.T) {
	ts.ClearTable("character_inventory_items")
	ts.ClearTable("characters")
	ts.ClearTable("armour")
	ts.ClearTable("items")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	characterOne := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, characterOne)

	characterTwo := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, characterTwo)

	characterThree := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, characterThree)

	itemOne := &models.Item{ID: 1, Name: "Test Item One"}
	factories.NewItem(ts.S.Db, itemOne)
	armourOne := &models.Armour{ItemID: 1}
	factories.NewArmour(ts.S.Db, armourOne)

	itemTwo := &models.Item{ID: 2, Name: "Test Item Two"}
	factories.NewItem(ts.S.Db, itemTwo)
	armourTwo := &models.Armour{ItemID: 2}
	factories.NewArmour(ts.S.Db, armourTwo)

	itemThree := &models.Item{ID: 3, Name: "Test Item Three"}
	factories.NewItem(ts.S.Db, itemThree)
	armourThree := &models.Armour{ItemID: 3}
	factories.NewArmour(ts.S.Db, armourThree)

	itemFour := &models.Item{ID: 4, Name: "Test Item Four"}
	factories.NewItem(ts.S.Db, itemFour)

	equippedArmour := &models.CharacterInventoryItem{CharacterID: characterOne.ID, ItemID: itemOne.ID, Equipped: true, Type: "armour"}
	factories.NewCharacterInventoryItem(ts.S.Db, equippedArmour)

	unequippedArmour := &models.CharacterInventoryItem{CharacterID: characterTwo.ID, ItemID: itemTwo.ID, Equipped: false, Type: "armour"}
	factories.NewCharacterInventoryItem(ts.S.Db, unequippedArmour)

	notArmour := &models.CharacterInventoryItem{CharacterID: characterOne.ID, ItemID: itemFour.ID}
	factories.NewCharacterInventoryItem(ts.S.Db, notArmour)

	characterTwoEquippedArmour := &models.CharacterInventoryItem{CharacterID: characterTwo.ID, ItemID: itemThree.ID, Equipped: true, Type: "armour"}
	factories.NewCharacterInventoryItem(ts.S.Db, characterTwoEquippedArmour)

	cases := []helpers.TestCase{
		{
			TestName: "Can get equipped weapons for character - not unequipped, non-weapons or other characters",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/inventory/equipped-armour",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%s"`, itemOne.Name),
					fmt.Sprintf(`"item_id":%v`, itemOne.ID),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%s"`, itemTwo.Name),
					fmt.Sprintf(`"name":"%s"`, itemThree.Name),
					fmt.Sprintf(`"name":"%s"`, itemFour.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestToggleCharacterInventoryItemEquipped(t *testing.T) {
	ts.ClearTable("character_inventory_items")
	ts.ClearTable("characters")
	ts.ClearTable("items")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	character := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, character)

	characterTwo := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, characterTwo)

	itemOne := &models.Item{ID: 1, Equippable: true}
	factories.NewItem(ts.S.Db, itemOne)
	weaponOne := &models.Weapon{ItemID: 1}
	factories.NewWeapon(ts.S.Db, weaponOne)

	itemTwo := &models.Item{ID: 2, Equippable: true}
	factories.NewItem(ts.S.Db, itemTwo)
	weaponTwo := &models.Weapon{ItemID: 2}
	factories.NewWeapon(ts.S.Db, weaponTwo)

	itemThree := &models.Item{ID: 3, Equippable: false}
	factories.NewItem(ts.S.Db, itemThree)

	itemFour := &models.Item{ID: 4, Equippable: true}
	factories.NewItem(ts.S.Db, itemFour)

	equipmentWeapon := &models.CharacterInventoryItem{
		ID:          1,
		CharacterID: character.ID,
		ItemID:      itemOne.ID,
		Equipped:    true,
		Location:    "Equipment",
	}
	factories.NewCharacterInventoryItem(ts.S.Db, equipmentWeapon)

	backpackWeapon := &models.CharacterInventoryItem{
		ID:          2,
		CharacterID: character.ID,
		ItemID:      itemTwo.ID,
		Equipped:    false,
		Location:    "Backpack",
	}
	factories.NewCharacterInventoryItem(ts.S.Db, backpackWeapon)

	equipmentItem := &models.CharacterInventoryItem{
		ID:          3,
		CharacterID: character.ID,
		ItemID:      itemThree.ID,
		Location:    "Equipment",
	}
	factories.NewCharacterInventoryItem(ts.S.Db, equipmentItem)

	otherCharacterEquipmentWeapon := &models.CharacterInventoryItem{
		ID:          4,
		CharacterID: characterTwo.ID,
		ItemID:      itemFour.ID,
		Equipped:    true,
		Location:    "Equipment",
	}
	factories.NewCharacterInventoryItem(ts.S.Db, otherCharacterEquipmentWeapon)

	cases := []helpers.TestCase{
		{
			TestName: "Can update character inventory item",
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("/characters/%v/inventory/%v", character.ID, equipmentWeapon.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"id":%v`, equipmentWeapon.ID),
					fmt.Sprintf(`"equipped":%v`, !equipmentWeapon.Equipped),
				},
			},
		},
		{
			TestName: "Can't equip items in character backpack",
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("/characters/%v/inventory/%v", character.ID, backpackWeapon.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"id":%v`, backpackWeapon.ID),
					fmt.Sprintf(`"equipped":%v`, backpackWeapon.Equipped),
				},
			},
		},
		{
			TestName: "Can't equip items that are unequippable",
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("/characters/%v/inventory/%v", character.ID, equipmentItem.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "404 Response for invalid character id",
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("/characters/invalid-id/inventory/%v", equipmentWeapon.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			TestName: "404 Response for invalid item id",
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("/characters/%v/inventory/invalid-id", character.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
