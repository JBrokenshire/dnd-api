package tests

import (
	"dnd-api/api/requests"
	"dnd-api/db/models"
	"dnd-api/services/jwt_service"
	"dnd-api/tests/helpers"
	"fmt"
	"net/http"
	"testing"
	"time"
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

func TestAuth_RefreshToken(t *testing.T) {
	ts.ClearTable("users")
	ts.SetupDefaultUsers()

	tokenService := jwt_service.NewTokenService()

	cookie := tokenService.GetRefreshCookie(ts.AdminRefreshToken, time.Now().Add(1*time.Hour))

	notExistUser := models.User{ID: 1000}
	notExistToken, notExistExpiry, err := tokenService.CreateUserRefreshToken(&notExistUser)
	if err != nil {
		t.Fatalf("Error creating refresh token: %v", err)
	}
	notExistCookie := tokenService.GetRefreshCookie(notExistToken, *notExistExpiry)

	invalidToken := ts.AdminToken[1 : len(ts.AdminRefreshToken)-1]
	invalidTokenCookie := tokenService.GetRefreshCookie(invalidToken, *notExistExpiry)

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/auth/refresh",
	}

	cases := []helpers.TestCase{
		{
			Name:           "Can't refresh token without providing refresh cookie",
			Request:        request,
			RequestCookies: []*http.Cookie{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "No refresh cookie provided",
			},
		},
		{
			Name:           "Can't refresh token with invalid refresh cookie",
			Request:        request,
			RequestCookies: []*http.Cookie{invalidTokenCookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyPart:   "Invalid token",
			},
		},
		{
			Name:           "Can't refresh token for user that doesn't exist",
			Request:        request,
			RequestCookies: []*http.Cookie{notExistCookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "User not found",
			},
		},
		{
			Name:           "Can refresh token",
			Request:        request,
			RequestCookies: []*http.Cookie{cookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
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

func TestAuth_Logout(t *testing.T) {
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/auth/logout",
	}

	tokenService := jwt_service.NewTokenService()

	adminRefreshCookie := tokenService.GetRefreshCookie(ts.AdminRefreshToken, time.Now().Add(1*time.Hour))

	cases := []helpers.TestCase{
		{
			Name:           "Can logout",
			Request:        request,
			RequestCookies: []*http.Cookie{adminRefreshCookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyParts: []string{
					"Logged Out",
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
