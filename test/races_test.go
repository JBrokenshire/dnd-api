package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAllRaces(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get list of races (populated table)",
			Setup: func() {
				ts.ClearTable("characters") // Have to clear characters first because of foreign key constraint
				ts.ClearTable("races")
				ts.SetupDefaultRaces()
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/races"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Aarakocra"`,
					`"name":"Stout Halfling"`,
					`"name":"Variant Aasimar"`,
				},
			},
		},
		{
			TestName: "can get list of races (empty table)",
			Setup: func() {
				ts.ClearTable("characters") // Have to clear characters first because of foreign key constraint
				ts.ClearTable("races")
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/races"},
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

func TestGetRace(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get race by id",
			Setup: func() {
				ts.SetupDefaultRaces()
			},
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/races/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Aarakocra"`},
			},
		},
		{
			TestName: "get /races/:id returns 404 not found on race id not in database",
			Setup: func() {
				ts.SetupDefaultRaces()
			},
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/races/100",
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

func TestGetRaceTraits(t *testing.T) {
	ts.ClearTable("race_traits")
	ts.ClearTable("races")
	ts.ClearTable("traits")

	raceOne := &models.Race{}
	factories.NewRace(ts.S.Db, raceOne)
	traitOne := &models.Trait{}
	factories.NewTrait(ts.S.Db, traitOne)
	raceTraitOne := &models.RaceTrait{RaceID: raceOne.ID, TraitID: traitOne.ID}
	factories.NewRaceTrait(ts.S.Db, raceTraitOne)

	noTraits := &models.Race{}
	factories.NewRace(ts.S.Db, noTraits)

	cases := []helpers.TestCase{
		{
			TestName: "Can get traits for race",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/races/%v/traits", raceOne.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   fmt.Sprintf(`"name":"%v"`, traitOne.Name),
			},
		},
		{
			TestName: "Can get empty response for race with no traits",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/races/%v/traits", noTraits.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "[]",
			},
		},
		{
			TestName: "Can get 404 for invalid race id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/races/invalid-id/traits",
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
