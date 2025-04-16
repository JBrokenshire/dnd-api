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

func TestCharacter_List(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.SetupDefaultUsers()

	// Create class
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create race
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	// Create Characters
	character := &m.Character{ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, character)
	character2 := &m.Character{ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, character2)
	namedCharacter := &m.Character{Name: "Test Character", ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, namedCharacter)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/characters%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get characters",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, differentUserCharacter.Name),
				},
			},
		},
		{
			Name:    "Can get page 0 of characters",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, differentUserCharacter.Name),
				},
			},
		},
		{
			Name:    "Can get page 1 of characters",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, differentUserCharacter.Name),
				},
			},
		},
		{
			Name:    "Can filter characters by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, differentUserCharacter.Name),
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

func TestCharacter_Get(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.SetupDefaultUsers()

	// Create class
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create race
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	// Create characters
	character := &m.Character{ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, character)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/characters/%v", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't get character that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't get character with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't get character that belongs to a different user",
			Request: getRequest(differentUserCharacter.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can get character",
			Request: getRequest(character.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
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

func TestCharacter_Create(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.SetupDefaultUsers()

	// Create class
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create race
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/characters",
	}

	RunNoAuthenticationTests(t, request.Method, request.Url)

	cases := []helpers.TestCase{
		{
			Name:        "Can't create character without required fields",
			Request:     request,
			RequestBody: requests.CreateCharacterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name is a required field",
					"ClassId is a required field",
					"RaceId is a required field",
				},
			},
		},
		{
			Name:    "Can't create character if fields exceed max length",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:    string(make([]byte, 201)),
				ClassId: class.ID,
				RaceId:  race.ID,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can't create character with class id",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:    "Test Name",
				ClassId: 1000,
				RaceId:  race.ID,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't create character with race id",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:    "Test Name",
				ClassId: class.ID,
				RaceId:  1000,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can create character",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:    "Test Name",
				ClassId: class.ID,
				RaceId:  race.ID,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, "Test Name"),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"user_id":%v`, ts.AdminUser.ID),
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character was created",
					Model: m.Character{
						UserId:  ts.AdminUser.ID,
						Name:    "Test Name",
						ClassId: class.ID,
						RaceId:  race.ID,
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

func TestCharacter_Update(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.SetupDefaultUsers()

	// Create classes
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)
	class2 := &m.Class{}
	factories.NewClass(ts.S.Db, class2)

	// Create races
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)
	race2 := &m.Race{}
	factories.NewRace(ts.S.Db, race2)

	// Create characters
	character := &m.Character{RaceId: race.ID, ClassId: class.ID}
	factories.NewCharacter(ts.S.Db, character)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			Url:    fmt.Sprintf("/characters/%v", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:        "Can't update character without required fields",
			Request:     getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name is a required field",
					"ClassId is a required field",
					"RaceId is a required field",
				},
			},
		},
		{
			Name:    "Can't update character if fields exceed max length",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:    string(make([]byte, 201)),
				ClassId: class.ID,
				RaceId:  race.ID,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can't update character that doesn't exist",
			Request: getRequest(1000),
			RequestBody: requests.UpdateCharacterRequest{
				Name:    "Test Name",
				ClassId: class.ID,
				RaceId:  race.ID,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't update character that belongs to a different user",
			Request: getRequest(differentUserCharacter.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:    "Test Name",
				ClassId: class.ID,
				RaceId:  race.ID,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't update character with class id that doesn't exist",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:    "Test Name",
				ClassId: 1000,
				RaceId:  race.ID,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't update character with race id that doesn't exist",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:    "Test Name",
				ClassId: class.ID,
				RaceId:  1000,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can update character",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:    "Test Name",
				ClassId: class2.ID,
				RaceId:  race2.ID,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Test Name"`,
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character was updated",
					Model: m.Character{
						ID:      character.ID,
						UserId:  ts.AdminUser.ID,
						Name:    "Test Name",
						ClassId: class2.ID,
						RaceId:  race2.ID,
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

func TestCharacter_Delete(t *testing.T) {
	ts.ClearTable("characters")
	ts.SetupDefaultUsers()

	// Create characters
	character := &m.Character{}
	factories.NewCharacter(ts.S.Db, character)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodDelete,
			Url:    fmt.Sprintf("/characters/%v", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't delete character that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't delete character with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't delete character that belongs to a different user",
			Request: getRequest(differentUserCharacter.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can delete character",
			Request: getRequest(character.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Character deleted successfully",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character was deleted",
					Model: m.Character{
						ID:   character.ID,
						Name: character.Name,
					},
					CountExpected: 0,
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
