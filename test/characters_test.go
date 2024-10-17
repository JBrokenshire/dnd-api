package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/server/requests"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAllCharacters(t *testing.T) {
	ts.SetupDefaultCharacters()

	cases := []helpers.TestCase{
		{
			TestName: "can get list of characters (populated table)",
			Request:  helpers.Request{Method: http.MethodGet, URL: "/characters"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Faelan Haversham"`,
					`"name":"PeeWee McAnkle-Biter"`,
				},
			},
		},
		{
			TestName: "can get list of characters (empty table)",
			Setup: func() {
				ts.ClearTable("characters")
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/characters"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "[]",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestCreateCharacter(t *testing.T) {
	ts.SetupDefaultCharacters()

	request := helpers.Request{
		Method: http.MethodPost,
		URL:    "/characters",
	}

	characterRequest := requests.NewCharacterRequest(&requests.CharacterRequest{})

	cases := []helpers.TestCase{
		{
			TestName:    "can create a new character from valid json in request body",
			Request:     request,
			RequestBody: characterRequest,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, characterRequest.Name),
					fmt.Sprintf(`"level":%v`, characterRequest.Level),
					fmt.Sprintf(`"class_id":%v`, characterRequest.ClassID),
					fmt.Sprintf(`"race_id":%v`, characterRequest.RaceID),
					fmt.Sprintf(`"profile_picture_url":%v`, characterRequest.ProfilePictureURL),
					fmt.Sprintf(`"strength":%v`, characterRequest.Strength),
					fmt.Sprintf(`"dexterity":%v`, characterRequest.Dexterity),
					fmt.Sprintf(`"constitution":%v`, characterRequest.Constitution),
					fmt.Sprintf(`"intelligence":%v`, characterRequest.Intelligence),
					fmt.Sprintf(`"wisdom":%v`, characterRequest.Wisdom),
					fmt.Sprintf(`"charisma":%v`, characterRequest.Charisma),
					fmt.Sprintf(`"proficient_strength":%v`, characterRequest.ProficientStrength),
					fmt.Sprintf(`"proficient_dexterity":%v`, characterRequest.ProficientDexterity),
					fmt.Sprintf(`"proficient_constitution":%v`, characterRequest.ProficientConstitution),
					fmt.Sprintf(`"proficient_intelligence":%v`, characterRequest.ProficientIntelligence),
					fmt.Sprintf(`"proficient_wisdom":%v`, characterRequest.ProficientWisdom),
					fmt.Sprintf(`"proficient_charisma":%v`, characterRequest.ProficientCharisma),
					fmt.Sprintf(`"walking_speed_modifier":%v`, characterRequest.WalkingSpeedModifier),
					fmt.Sprintf(`"inspiration":%v`, characterRequest.Inspiration),
					fmt.Sprintf(`"current_hit_points":%v`, characterRequest.CurrentHitPoints),
					fmt.Sprintf(`"max_hit_points":%v`, characterRequest.MaxHitPoints),
					fmt.Sprintf(`"temp_hit_points":%v`, characterRequest.TempHitPoints),
					fmt.Sprintf(`"initiative_modifier":%v`, characterRequest.InitiativeModifier),
					fmt.Sprintf(`"attacks_per_action":%v`, characterRequest.AttacksPerAction),
					fmt.Sprintf(`"background_name":"%v"`, characterRequest.BackgroundName),
				},
			},
		},
		{
			TestName: "post /character/:id 400 bad request on no request body",
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid request body",
			},
		},
		{
			TestName: "post /character/:id 400 internal server error on invalid class id",
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "test",
				Level:   1,
				ClassID: 1000,
				RaceID:  1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character classID",
			},
		},
		{
			TestName: "post /character/:id 400 bad request on no character name",
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Level:   1,
				ClassID: 1,
				RaceID:  1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character name",
			},
		},
		{
			TestName: "post /character/:id 400 internal server error on invalid race id",
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "test",
				Level:   1,
				ClassID: 1,
				RaceID:  1000,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character raceID",
			},
		},
		{
			TestName: "post /character/:id 400 internal server error on invalid level",
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "test",
				Level:   1000,
				ClassID: 1,
				RaceID:  1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character level",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestGetCharacter(t *testing.T) {
	ts.SetupDefaultCharacters()

	cases := []helpers.TestCase{
		{
			TestName: "can get character by id",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Faelan Haversham"`, `"level":3`, `"class_id":3`, `"race_id":18`},
			},
		},
		{
			TestName: "get /characters/:id returns 404 not found on character id not in database",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/10",
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

func TestUpdateCharacter(t *testing.T) {
	ts.SetupDefaultCharacters()

	request := helpers.Request{
		Method: http.MethodPut,
		URL:    "/characters/1",
	}

	cases := []helpers.TestCase{
		{
			TestName: "can update character by id with valid json in request body",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			RequestBody: requests.CharacterRequest{Name: "Test",
				ClassID: 1,
				RaceID:  1,
				Level:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Test"`, `"level":1`, `"class_id":1`, `"race_id":1`},
			},
		},
		{
			TestName: "put /characters/:id returns 404 not found on character id not in database",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/characters/10",
			},
			RequestBody: requests.CharacterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			TestName: "put /characters/:id returns 400 bad request on empty request body",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character request body",
			},
		},
		{
			TestName:    "no update on empty character in request body",
			Setup:       ts.SetupDefaultCharacters,
			Request:     request,
			RequestBody: requests.CharacterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Faelan Haversham"`, `"level":3`, `"class_id":3`, `"race_id":18`},
			},
		},
		{
			TestName: "update with only name in request body",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name: "Test",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Test"`, `"level":3`, `"class_id":3`, `"race_id":18`},
			},
		},
		{
			TestName: "update with only class id in request body",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			RequestBody: requests.CharacterRequest{
				ClassID: 1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Faelan Haversham"`, `"level":3`, `"class_id":1`, `"race_id":18`},
			},
		},
		{
			TestName: "put /character/:id returns 400 bad request on invalid class id",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/characters/1",
			},
			RequestBody: requests.CharacterRequest{
				Name:    "Test",
				ClassID: 1000,
				RaceID:  1,
				Level:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character classID",
			},
		},
		{
			TestName: "put /character/:id returns 400 bad request on invalid race id",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/characters/1",
			},
			RequestBody: requests.CharacterRequest{
				Name:    "Test",
				ClassID: 1,
				RaceID:  1000,
				Level:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character raceID",
			},
		},
		{
			TestName: "put /characters/:id returns 400 bad request on invalid level",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "Test",
				ClassID: 1,
				RaceID:  1,
				Level:   100,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character level",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestCharacterToggleInspiration(t *testing.T) {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	withInspiration := &models.Character{
		Inspiration: true,
	}
	factories.NewCharacter(ts.S.Db, withInspiration)

	withoutInspiration := &models.Character{
		Inspiration: false,
	}
	factories.NewCharacter(ts.S.Db, withoutInspiration)

	cases := []helpers.TestCase{
		{
			TestName: "can change character inspiration from true to false",
			Request: helpers.Request{
				Method: "GET",
				URL:    fmt.Sprintf("/characters/%v/inspiration", withInspiration.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, withInspiration.Name),
					fmt.Sprintf(`"inspiration":%v`, false),
				},
			},
		},
		{
			TestName: "can change character inspiration from true to false",
			Request: helpers.Request{
				Method: "GET",
				URL:    fmt.Sprintf("/characters/%v/inspiration", withoutInspiration.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, withoutInspiration.Name),
					fmt.Sprintf(`"inspiration":%v`, true),
				},
			},
		},
		{
			TestName: "404 not found on invalid character id",
			Request: helpers.Request{
				Method: "GET",
				URL:    fmt.Sprintf("/characters/invalid-id/inspiration"),
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

func TestLevelUpCharacter(t *testing.T) {
	ts.SetupDefaultCharacters()

	cases := []helpers.TestCase{
		{
			TestName: "can level up character by id",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/level-up",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Faelan Haversham"`, `"level":4`, `"class_id":3`, `"race_id":18`},
			},
		},
		{
			TestName: "put /characters/:id/level-up returns 404 not found on character id not in database",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/10/level-up",
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

func TestUpdateCharacterHealth(t *testing.T) {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	healCharacter := &models.Character{
		CurrentHitPoints: 10,
		MaxHitPoints:     20,
	}
	factories.NewCharacter(ts.S.Db, healCharacter)
	damageCharacter := &models.Character{
		CurrentHitPoints: 10,
		MaxHitPoints:     20,
	}
	factories.NewCharacter(ts.S.Db, damageCharacter)

	cases := []helpers.TestCase{
		{
			TestName: "Can heal character 1 hit point",
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("/characters/%v/heal/1", healCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, healCharacter.Name),
					fmt.Sprintf(`"current_hit_points":%v`, 11),
				},
			},
		},
		{
			TestName: "Character can't heal more than max hit points",
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("/characters/%v/heal/1000", healCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, healCharacter.Name),
					fmt.Sprintf(`"current_hit_points":%v`, healCharacter.MaxHitPoints),
				},
			},
		},
		{
			TestName: "Can damage character 1 hit point",
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("/characters/%v/damage/1", damageCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, damageCharacter.Name),
					fmt.Sprintf(`"current_hit_points":%v`, 9),
				},
			},
		},
		{
			TestName: "Character can't go below 0 hit points",
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    fmt.Sprintf("/characters/%v/damage/1000", damageCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, damageCharacter.Name),
					fmt.Sprintf(`"current_hit_points":%v`, 0),
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

func TestCharacterArmourClass(t *testing.T) {
	ts.ClearTable("characters")
	ts.ClearTable("character_inventory_items")
	ts.ClearTable("armour")
	ts.ClearTable("shields")
	ts.ClearTable("items")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	unarmouredCharacter := models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, &unarmouredCharacter)

	dexterityCharacter := models.Character{ID: 2, Dexterity: 20}
	factories.NewCharacter(ts.S.Db, &dexterityCharacter)

	armourItem := models.Item{ID: 1, Equippable: true}
	factories.NewItem(ts.S.Db, &armourItem)
	armour := models.Armour{ItemID: 1, BaseAC: 12, MaxDexterityModifier: 10}
	factories.NewArmour(ts.S.Db, &armour)
	armouredCharacter := models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, &armouredCharacter)
	armouredCharacterInventoryItem := models.CharacterInventoryItem{CharacterID: armouredCharacter.ID, ItemID: armourItem.ID, Equipped: true, Location: "Equipment", Type: "armour"}
	factories.NewCharacterInventoryItem(ts.S.Db, &armouredCharacterInventoryItem)

	shieldItem := models.Item{ID: 2, Equippable: true}
	factories.NewItem(ts.S.Db, &shieldItem)
	shield := models.Shield{ItemID: 2, BonusAC: 2}
	factories.NewShield(ts.S.Db, &shield)
	shieldCharacter := models.Character{ID: 4, Dexterity: 10}
	factories.NewCharacter(ts.S.Db, &shieldCharacter)
	shieldCharacterInventoryItem := models.CharacterInventoryItem{
		CharacterID: shieldCharacter.ID,
		ItemID:      shieldItem.ID,
		Type:        "shield",
		Equipped:    true,
		Location:    "Equipment",
	}
	factories.NewCharacterInventoryItem(ts.S.Db, &shieldCharacterInventoryItem)

	shieldArmourCharacter := models.Character{ID: 5, Dexterity: 10}
	factories.NewCharacter(ts.S.Db, &shieldArmourCharacter)
	shieldArmourCharacterArmour := models.CharacterInventoryItem{
		CharacterID: shieldArmourCharacter.ID,
		ItemID:      armourItem.ID,
		Type:        "armour",
		Equipped:    true,
		Location:    "Equipment",
	}
	factories.NewCharacterInventoryItem(ts.S.Db, &shieldArmourCharacterArmour)
	shieldArmourCharacterShield := models.CharacterInventoryItem{
		CharacterID: shieldArmourCharacter.ID,
		ItemID:      shieldItem.ID,
		Type:        "shield",
		Equipped:    true,
		Location:    "Equipment",
	}
	factories.NewCharacterInventoryItem(ts.S.Db, &shieldArmourCharacterShield)

	maxDexterityItem := models.Item{ID: 3, Equippable: true}
	factories.NewItem(ts.S.Db, &maxDexterityItem)
	maxDexterityArmour := models.Armour{
		ItemID:               3,
		BaseAC:               12,
		MaxDexterityModifier: 2,
	}
	factories.NewArmour(ts.S.Db, &maxDexterityArmour)
	maxDexterityCharacter := models.Character{ID: 6, Dexterity: 20}
	factories.NewCharacter(ts.S.Db, &maxDexterityCharacter)
	maxDexterityCharacterArmour := models.CharacterInventoryItem{
		CharacterID: maxDexterityCharacter.ID,
		ItemID:      maxDexterityItem.ID,
		Type:        "armour",
		Equipped:    true,
		Location:    "Equipment",
	}
	factories.NewCharacterInventoryItem(ts.S.Db, &maxDexterityCharacterArmour)

	cases := []helpers.TestCase{
		{
			TestName: "Can get armour class for character with no armour and no shield",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/armour-class", unarmouredCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "10",
			},
		},
		{
			TestName: "Can calculate armour class based on character dexterity",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/armour-class", dexterityCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "15",
			},
		},
		{
			TestName: "Can get armour class for character with armour",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/armour-class", armouredCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "12",
			},
		},
		{
			TestName: "Can get armour class for character with shield",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/armour-class", shieldCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "12",
			},
		},
		{
			TestName: "Can get armour class for character with armour and shield",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/armour-class", shieldArmourCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "14",
			},
		},
		{
			TestName: "Can get armour class for character with dexterity higher than armour max dexterity",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/armour-class", maxDexterityCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "14",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestDeleteCharacter(t *testing.T) {
	ts.SetupDefaultCharacters()

	cases := []helpers.TestCase{
		{
			TestName: "can delete character by id",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodDelete,
				URL:    "/characters/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "character successfully deleted",
			},
		},
		{
			TestName: "delete /characters/:id returns 404 not found on character id not in database",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodDelete,
				URL:    "/characters/10",
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
