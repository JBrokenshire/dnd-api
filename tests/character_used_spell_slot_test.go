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

func TestCharacterUsedSpellSlot_Update(t *testing.T) {
	ts.ClearTable("characters")
	ts.ClearTable("character_used_spell_slots")
	ts.SetupDefaultUsers()

	// Create characters
	character := &m.Character{}
	factories.NewCharacter(ts.S.Db, character)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	// Create character used spell slots
	characterUsedSpellSlot := &m.CharacterUsedSpellSlot{CharacterId: character.ID, SpellLevel: 1}
	factories.NewCharacterUsedSpellSlot(ts.S.Db, characterUsedSpellSlot)

	getRequest := func(characterID, spellLevel interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			Url:    fmt.Sprintf("/characters/%v/spell-slots/%v", characterID, spellLevel),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't update character used spell slots with invalid request body",
			Request:     getRequest(character.ID, 1),
			RequestBody: requests.UpdateCharacterUsedSpellSlotRequest{SpellSlotsUsed: -1},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"SpellSlotsUsed must be 0 or greater",
				},
			},
		},
		{
			Name:        "Can't update character used spell slots with invalid spell level (string)",
			Request:     getRequest(character.ID, "invalid-spell-level"),
			RequestBody: requests.UpdateCharacterUsedSpellSlotRequest{SpellSlotsUsed: 1},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Spell level must be an integer",
			},
		},
		{
			Name:        "Can't update character used spell slots with invalid spell level (over 9)",
			Request:     getRequest(character.ID, 10),
			RequestBody: requests.UpdateCharacterUsedSpellSlotRequest{SpellSlotsUsed: 1},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Spell level must be between 1 and 9",
			},
		},
		{
			Name:        "Can't update character used spell slots with invalid spell level (under 1)",
			Request:     getRequest(character.ID, 0),
			RequestBody: requests.UpdateCharacterUsedSpellSlotRequest{SpellSlotsUsed: 1},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Spell level must be between 1 and 9",
			},
		},
		{
			Name:        "Can't update character used spell slots for character that doesn't exist",
			Request:     getRequest(1000, 1),
			RequestBody: requests.UpdateCharacterUsedSpellSlotRequest{SpellSlotsUsed: 1},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:        "Can't update character used spell slots for character that belongs to a different user",
			Request:     getRequest(differentUserCharacter.ID, 1),
			RequestBody: requests.UpdateCharacterUsedSpellSlotRequest{SpellSlotsUsed: 1},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:        "Can update character used spell slots",
			Request:     getRequest(character.ID, 1),
			RequestBody: requests.UpdateCharacterUsedSpellSlotRequest{SpellSlotsUsed: 1},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Character used spell slots updated",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character used spell slots updated",
					Model: m.CharacterUsedSpellSlot{
						CharacterId:    character.ID,
						SpellLevel:     1,
						SpellSlotsUsed: 1,
					},
					CountExpected: 1,
				},
			},
		},
		{
			Name:        "Can create new entry for used spell slots that doesn't already exist",
			Request:     getRequest(character.ID, 2),
			RequestBody: requests.UpdateCharacterUsedSpellSlotRequest{SpellSlotsUsed: 1},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Character used spell slots updated",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character used spell slots entry created",
					Model: m.CharacterUsedSpellSlot{
						CharacterId:    character.ID,
						SpellLevel:     2,
						SpellSlotsUsed: 1,
					},
					CountExpected: 1,
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
