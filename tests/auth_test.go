package tests

import (
	"dnd-api/api/requests"
	"dnd-api/tests/helpers"
	"net/http"
	"testing"
)

func TestAuth_Login(t *testing.T) {
	ts.ClearTable("users")
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/login",
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't login without the required fields",
			Request:     request,
			RequestBody: requests.LoginRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Username is a required field",
					"Password is a required field",
				},
			},
		},
		{
			Name:    "Can't login with username that doesn't exist",
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: "Invalid Username",
				Password: "Invalid Password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "User not found",
			},
		},
		{
			Name:    "Can't login with incorrect password",
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: helpers.AdminUsername,
				Password: "Incorrect Password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyPart:   "Incorrect password. Please try again.",
			},
		},
		{
			Name:    "Can login with correct credentials",
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: helpers.AdminUsername,
				Password: helpers.AdminPassword,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"authorised":true`,
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
