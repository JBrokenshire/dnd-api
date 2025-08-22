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

func TestCharacterHealth_Update(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")

	// Create class
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create race
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	// Create characters
	character := &m.Character{ClassId: class.ID, RaceId: race.ID, MaxHitPoints: 100}
	factories.NewCharacter(ts.S.Db, character)
	differentUserCharacter := &m.Character{UserId: 1000, ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			Url:    fmt.Sprintf("/characters/%v/health", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't update health with value less than 0",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterHealthRequest{
				CurrentHitPoints: -1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"CurrentHitPoints must be 0 or greater",
				},
			},
		},
		{
			Name:    "Can't update health for character that doesn't exist",
			Request: getRequest(1000),
			RequestBody: requests.UpdateCharacterHealthRequest{
				CurrentHitPoints: 10,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't update health for character that belongs to a different user",
			Request: getRequest(differentUserCharacter.ID),
			RequestBody: requests.UpdateCharacterHealthRequest{
				CurrentHitPoints: 10,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can update health for character",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterHealthRequest{
				CurrentHitPoints: 10,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					`"current_hit_points":10`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character was updated",
					Model: m.Character{
						ID:               character.ID,
						CurrentHitPoints: 10,
					},
					CountExpected: 1,
				},
			},
		},
		{
			Name:    "Can update character current only as high as max hit points",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterHealthRequest{
				CurrentHitPoints: 101,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"id":%v`, character.ID),
					`"current_hit_points":100`,
				},
				BodyPartsMissing: []string{
					`"current_hit_points":101`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character was updated",
					Model: m.Character{
						ID:               character.ID,
						CurrentHitPoints: 100,
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
