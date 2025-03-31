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

func TestRace_List(t *testing.T) {
	ts.ClearTable("races")

	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)
	race2 := &m.Race{}
	factories.NewRace(ts.S.Db, race2)
	race3 := &m.Race{}
	factories.NewRace(ts.S.Db, race3)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			URL:    fmt.Sprintf("/races%v", query),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:    "Can get races",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, race3.Name),
					`"total_count":3`,
				},
			},
		},
		{
			Name:    "Can get page 0 of races",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, race3.Name)},
			},
		},
		{
			Name:    "Can get page 1 of races",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, race3.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestRace_Get(t *testing.T) {
	ts.ClearTable("races")

	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			URL:    fmt.Sprintf("/races/%v", id),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:    "Can't get race that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't get race with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can get race",
			Request: getRequest(race.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestRace_Create(t *testing.T) {
	ts.ClearTable("races")

	request := helpers.Request{
		Method: http.MethodPost,
		URL:    "/races",
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't create race without required fields",
			Request:     request,
			RequestBody: &requests.CreateRaceRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Name is a required field",
				},
			},
		},
		{
			Name:    "Can't create race if fields exceed max length",
			Request: request,
			RequestBody: &requests.CreateRaceRequest{
				Name: string(make([]byte, 201)),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Name must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can create race",
			Request: request,
			RequestBody: &requests.CreateRaceRequest{
				Name: "Test Race",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					`"name":"Test Race"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Race was created",
					Model: m.Race{
						Name: "Test Race",
					},
					CountExpected: 1,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestRace_Update(t *testing.T) {
	ts.ClearTable("races")

	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			URL:    fmt.Sprintf("/races/%v", id),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't update race without required fields",
			Request:     getRequest(race.ID),
			RequestBody: requests.UpdateRaceRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Name is a required field",
				},
			},
		},
		{
			Name:    "Can't update race if fields exceed max length",
			Request: getRequest(race.ID),
			RequestBody: requests.UpdateRaceRequest{
				Name: string(make([]byte, 201)),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Name must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can't update race that doesn't exist",
			Request: getRequest(1000),
			RequestBody: requests.UpdateRaceRequest{
				Name: "Updated Name",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't update race with invalid id",
			Request: getRequest("invalid-id"),
			RequestBody: requests.UpdateRaceRequest{
				Name: "Updated Name",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can update race",
			Request: getRequest(race.ID),
			RequestBody: requests.UpdateRaceRequest{
				Name: "Updated Name",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Updated Name"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Race was updated",
					Model: m.Race{
						Name: "Updated Name",
					},
					CountExpected: 1,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestRace_Delete(t *testing.T) {
	ts.ClearTable("races")

	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodDelete,
			URL:    fmt.Sprintf("/races/%v", id),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:    "Can't delete race that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't delete race with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can delete race",
			Request: getRequest(race.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Race deleted successfully",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Race was deleted",
					Model: m.Race{
						Name: race.Name,
					},
					CountExpected: 0,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
