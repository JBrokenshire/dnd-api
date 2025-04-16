package tests

import (
	"dnd-api/api/requests"
	"dnd-api/db/factories"
	m "dnd-api/db/models"
	"dnd-api/tests/helpers"
	"net/http"
	"testing"
)

func TestUser_Create(t *testing.T) {
	ts.ClearTable("users")
	ts.SetupDefaultUsers()

	existingUser := &m.User{Username: "Test"}
	factories.NewUser(ts.S.Db, existingUser)

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/users",
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't create user without required fields",
			Request:     request,
			RequestBody: requests.CreateUserRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Username is a required field",
					"Password is a required field",
					"ConfirmPassword is a required field",
				},
			},
		},
		{
			Name:    "Can't create user if fields exceed max length",
			Request: request,
			RequestBody: requests.CreateUserRequest{
				Username:        string(make([]byte, 201)),
				Password:        string(make([]byte, 73)),
				ConfirmPassword: string(make([]byte, 73)),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Username must be a maximum of 200 characters in length",
					"Password must be a maximum of 72 characters in length",
					"ConfirmPassword must be a maximum of 72 characters in length",
				},
			},
		},
		{
			Name:    "Can't create user if provided passwords don't match",
			Request: request,
			RequestBody: requests.CreateUserRequest{
				Username:        "Test Username",
				Password:        "Test Password",
				ConfirmPassword: "Different Password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Passwords do not match",
			},
		},
		{
			Name:    "Can't create user with username that is already being used",
			Request: request,
			RequestBody: requests.CreateUserRequest{
				Username:        existingUser.Username,
				Password:        "Test Password",
				ConfirmPassword: "Test Password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "That username is already in use",
			},
		},
		{
			Name:    "Can create user",
			Request: request,
			RequestBody: requests.CreateUserRequest{
				Username:        "Test Username",
				Password:        "Test Password",
				ConfirmPassword: "Test Password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					`"username":"Test Username"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "User was created",
					Model: m.User{
						Username: "Test Username",
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
