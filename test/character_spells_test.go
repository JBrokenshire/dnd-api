package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCharacterHasSpells(t *testing.T) {
	ts.ClearTable("character_spells")
	ts.ClearTable("spells")
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()
	ts.SetupDefaultBackgrounds()

	hasSpellsCharacter := &models.Character{}
	factories.NewCharacter(ts.S.Db, hasSpellsCharacter)
	spell := &models.Spell{}
	factories.NewSpell(ts.S.Db, spell)
	characterSpell := &models.CharacterSpell{CharacterID: hasSpellsCharacter.ID, SpellID: hasSpellsCharacter.ID}
	factories.NewCharacterSpell(ts.S.Db, characterSpell)

	noSpellsCharacter := &models.Character{}
	factories.NewCharacter(ts.S.Db, noSpellsCharacter)

	cases := []helpers.TestCase{
		{
			TestName: "Can get true for character that has spells",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/has-spells", hasSpellsCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "true",
			},
		},
		{
			TestName: "Can get false for character with no spells",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/has-spells", noSpellsCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "false",
			},
		},
		{
			TestName: "Can get 404 for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/has-spells",
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

func TestGetCharacterSpells(t *testing.T) {
	ts.ClearTable("character_spells")
	ts.ClearTable("spells")
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()
	ts.SetupDefaultBackgrounds()

	characterOne := &models.Character{}
	factories.NewCharacter(ts.S.Db, characterOne)
	spellOne := &models.Spell{}
	factories.NewSpell(ts.S.Db, spellOne)
	characterOneSpell := &models.CharacterSpell{CharacterID: characterOne.ID, SpellID: spellOne.ID}
	factories.NewCharacterSpell(ts.S.Db, characterOneSpell)

	characterTwo := &models.Character{}
	factories.NewCharacter(ts.S.Db, characterTwo)
	spellTwo := &models.Spell{}
	factories.NewSpell(ts.S.Db, spellTwo)
	characterTwoSpell := &models.CharacterSpell{CharacterID: characterTwo.ID, SpellID: spellTwo.ID}
	factories.NewCharacterSpell(ts.S.Db, characterTwoSpell)

	noSpellsCharacter := &models.Character{}
	factories.NewCharacter(ts.S.Db, noSpellsCharacter)

	cases := []helpers.TestCase{
		{
			TestName: "Can get spells for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/spells", characterOne.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, spellOne.Name),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, spellTwo.Name),
				},
			},
		},
		{
			TestName: "Can get empty slice for character with no spells",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/spells", noSpellsCharacter.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   `[]`,
			},
		},
		{
			TestName: "Can get 404 for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/spells",
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
