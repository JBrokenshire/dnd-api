package tests

import (
	"dnd-api/db/factories"
	m "dnd-api/db/models"
	"dnd-api/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestCharacterInspiration_Update(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")

	// Create class
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create race
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	// Create character
	character := &m.Character{ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, character)
	differentUserCharacter := &m.Character{UserId: 1000, ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			Url:    fmt.Sprintf("/characters/%v/inspiration", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't toggle inspiration for character that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't toggle inspiration for character that belongs to a different user",
			Request: getRequest(differentUserCharacter.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can toggle inspiration for character",
			Request: getRequest(character.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"inspiration":true`),
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character was updated",
					Model: m.Character{
						ID:          character.ID,
						Inspiration: true,
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
