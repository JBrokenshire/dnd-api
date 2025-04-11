package tests

import (
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/pkg/jwt_service"
	"dnd-api/tests/helpers"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

// These tests run against the main test enterprise which has two factor turned on. They should only return an unauthenticated
// token and flags about the user (2fa set up or not)

// Run through various tests for logging in.
func TestAuthTableTest(t *testing.T) {

	// Clear the users table
	ts.ClearTable("users")
	ts.ClearTable("enterprises")
	ts.ClearTable("failed_logins")

	// Create an admin user so that we have some basic credentials in the database
	ts.CreateAdminUser()

	defaultUser := &models.User{
		Password: helpers.DefaultPasswordHash, // Abcd1234$
	}

	// Create the request
	request := helpers.Request{
		Method: http.MethodPost,
		URL:    "/auth/login",
	}

	cases := []helpers.TestCase{
		{
			Name: "Auth Success",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    helpers.AdminEmail,
				Password: helpers.AdminPassword,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyParts:  []string{`"authorised":false`},
				BodyPartsMissing: []string{
					"Made Purple", //Enterprise Name
					ts.AdminUser.Name,
					ts.AdminUser.Uid,
					"Group", // A group permission should be returned
				},
			},
		},
		{
			TestName: "Login attempt with incorrect password",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    "matt@purplevisits.com",
				Password: "incorrectPassword",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "Incorrect username or Password. Please try again.",
				DatabaseCheck: &helpers.DatabaseCheck{
					Model:         models.FailedLogin{Email: "matt@purplevisits.com", IpAddress: "127.0.0.0"},
					CountExpected: 1,
				},
			},
		},
		{
			TestName: "Login attempt as non-existent user",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    "user.not.exists@test.com",
				Password: "password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "Incorrect username or Password. Please try again.",
				DatabaseCheck: &helpers.DatabaseCheck{
					Model:         models.FailedLogin{Email: "user.not.exists@test.com", IpAddress: "127.0.0.0"},
					CountExpected: 1,
				},
			},
		},
		{
			TestName: "Login attempt with invalid enterprise",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    userInvalidEnterprise.Email,
				Password: "Abcd1234$",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "There is a problem with the Enterprise you are assigned to",
			},
		},
		{
			TestName: "Email is required",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Password: "password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "Email is a required field",
			},
		},
		{
			TestName: "Login attempt with no password",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email: "user.not.exists@test.com",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "Password is a required field",
			},
		},
		{
			TestName: "Login attempt with 3 previous failed password attempts",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    users[0].Email,
				Password: "Abcd1234$",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "Too many attempts, your account has been blocked temporarily.",
			},
		},
		{
			TestName: "Login attempt with 3 previous failed password attempts one from a different ip (wrong password)",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    users[1].Email,
				Password: "Abcd1234$5",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "Incorrect username or Password. Please try again.",
			},
		},
		{
			TestName: "Login attempt with 3 previous failed password attempts one from 15 minutes ago (wrong password)",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    users[2].Email,
				Password: "Abcd1234$5",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "Incorrect username or Password. Please try again.",
			},
		},
		{
			TestName: "Login attempt with 20 previous failed password attempts from 29 days ago (wrong password)",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    users[3].Email,
				Password: "Abcd1234$5",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "Too many attempts in the past month, your account has been blocked temporarily.",
			},
		},
		{
			TestName: "Login attempt with 20 previous failed password attempts from 31 days ago (wrong password)",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    users[4].Email,
				Password: "Abcd1234$5",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "Incorrect username or Password. Please try again.",
			},
		},
		{
			TestName: "Bad user agent should be blocked",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    helpers.AdminEmail,
				Password: helpers.AdminPassword,
			},
			RequestHeaders: map[string]string{"User-Agent": "Nessus"},
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "Error processing request (8492)",
			},
		},
		{
			TestName: "request is rejected if any of the fields exceeds the max amount of characters",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    fmt.Sprintf("a@%v.com", strings.Repeat("A", 195)),
				Password: string(make([]byte, 473)),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Email must be a maximum of 200 characters in length",
					"Password must be a maximum of 472 characters in length",
				},
			},
		},
		helpers.AddBadContentTestCase(request),
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {

			res := ts.ExecuteTestCase(&test)

			ValidateResults(t, test, res)

			//log.Println(res.Body.String())
			if res.Code == http.StatusOK {

				var authResponse responses.LoginResponse
				_ = json.Unmarshal([]byte(res.Body.String()), &authResponse)

				testToken(t, authResponse.AccessToken, helpers.AdminUid)

				cookies := res.Result().Cookies()
				if len(cookies) > 0 {
					t.Fatalf("Cookies were returned. 2FA should only return refresh token on second stage auth")
				}
			}
		})
	}
}

func TestAuthNo2FATest(t *testing.T) {

	// Clear the users table
	ts.ClearTable("users")
	ts.ClearTable("enterprises")
	ts.ClearTable("failed_logins")

	// Create an admin user so that we have some basic credentials in the database
	helpers.TestEnterprise.TwoFactorEnabled = false
	ts.CreateAdminUser()

	userNo2FA := &models.User{
		Password:         helpers.DefaultPasswordHash, // Abcd1234$
		TwoFactorEnabled: false,
	}

	factories.NewUser(ts.S.Db, userNo2FA)

	// Create the request
	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/login",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Auth Success with no 2FA",
			Request:  request,
			RequestBody: requests.LoginRequest{
				Email:    userNo2FA.Email,
				Password: "Abcd1234$",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyParts: []string{
					`"authorised":true`},
				BodyPartsMissing: []string{
					"Made Purple", //Enterprise Name
					userNo2FA.Name,
					userNo2FA.Uid,
					"Group", // A group permission should be returned
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {

			res := ts.ExecuteTestCase(&test)

			ValidateResults(t, test, res)

			//log.Println(res.Body.String())
			if res.Code == http.StatusOK {

				var authResponse responses.LoginResponse
				_ = json.Unmarshal([]byte(res.Body.String()), &authResponse)

				testToken(t, authResponse.AccessToken, userNo2FA.Uid)

				cookies := res.Result().Cookies()
				if len(cookies) == 0 {
					t.Fatalf("No Cookies were returned. 2FA is off so we should receive cookie!")
				}
			}
		})
	}
}

func TestAuthIpRangeTest(t *testing.T) {
	ts.ClearTable("failed_logins")
	ts.SetupDefaultUsers()

	//New Enterprise with user ip ranges
	enterpriseWithRanges := &models.Enterprise{
		UserIpRanges: "127.0.0.0/30",
	}
	factories.NewEnterpriseWithDefaults(ts.S.Db, enterpriseWithRanges)

	user := &models.User{
		Password:         helpers.DefaultPasswordHash, // Abcd1234$
		EnterpriseUID:    enterpriseWithRanges.Uid,
		SuperAdmin:       true,
		TwoFactorCode:    "VHX5EJYLRITO7VKMGCLUK6YOIXUC5EYP",
		TwoFactorEnabled: true,
	}

	factories.NewUser(ts.S.Db, user)

	// Create the request
	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/login",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Auth Success with valid Ip address",
			Request:  request,
			RequestHeaders: map[string]string{
				"X-Forwarded-For": "127.0.0.1",
			},
			RequestBody: requests.LoginRequest{
				Email:    user.Email,
				Password: "Abcd1234$",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyParts: []string{
					`"accessToken":`,
				},
			},
		},
		{
			TestName: "Cannot login with ip not in enterprise range",
			Request:  request,
			RequestHeaders: map[string]string{
				"X-Forwarded-For": "126.0.0.50",
			},
			RequestBody: requests.LoginRequest{
				Email:    user.Email,
				Password: "Abcd1234$",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyParts: []string{
					`Ip address is not allowed by enterprise`,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {

			res := ts.ExecuteTestCase(&test)

			ValidateResults(t, test, res)

			//log.Println(res.Body.String())
			if res.Code == http.StatusOK {

				var authResponse responses.LoginResponse
				_ = json.Unmarshal([]byte(res.Body.String()), &authResponse)

				testToken(t, authResponse.AccessToken, user.Uid)

				cookies := res.Result().Cookies()
				if len(cookies) > 0 {
					t.Fatalf("Cookies were returned. 2FA should only return refresh token on second stage auth")
				}
			}
		})
	}
}

func TestRefreshToken(t *testing.T) {

	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/auth/refresh",
	}

	tokenService := jwt_service.NewTokenService()

	// Generate good token for Admin
	adminRefreshCookie := tokenService.GetRefreshCookie(ts.AdminRefreshToken, time.Now().Add(1*time.Hour))

	notExistUser := models.User{Email: "user.not.exists@test.com"}
	notExistUser.ID = helpers.UserId + 1
	notExistToken, notExistExpiry, err := tokenService.CreateUserRefreshToken(&notExistUser, notExistUser.EnterpriseUID, ts.Enterprise.RefreshTokenDuration)
	if err != nil {
		t.Fatalf("Error creating refresh token: %v", err)
	}
	notExistCookie := tokenService.GetRefreshCookie(notExistToken, *notExistExpiry)

	invalidToken := ts.AdminToken[1 : len(ts.AdminRefreshToken)-1]
	invalidTokenCookie := tokenService.GetRefreshCookie(invalidToken, *notExistExpiry)

	cases := []helpers.TestCase{
		{
			TestName:       "Refresh success",
			Request:        request,
			RequestCookies: []*http.Cookie{adminRefreshCookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyParts: []string{
					"Made Purple", // The enterprise Name
					ts.AdminUser.Name,
					ts.AdminUser.Uid,
					"Group", // A group permission should be returned
					`"mdm_version":2`,
				},
			},
		},
		{
			TestName:       "Refresh token of non-existent user",
			Request:        request,
			RequestCookies: []*http.Cookie{notExistCookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "User not found",
			},
		},
		{
			TestName:       "Refresh invalid token",
			Request:        request,
			RequestCookies: []*http.Cookie{invalidTokenCookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "error",
			},
		},
		{
			TestName: "Cookie is required",
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "No refresh cookie provided",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			res := ts.ExecuteTestCase(&test)

			// Validate the results from the test case
			ValidateResults(t, test, res)

			if res.Code == http.StatusOK {
				var authResponse responses.RefreshResponse
				_ = json.Unmarshal([]byte(res.Body.String()), &authResponse)

				testToken(t, authResponse.AccessToken, helpers.AdminUid)

				// Test the refresh token in the cookie
				cookies := res.Result().Cookies()
				if len(cookies) == 0 {
					t.Fatalf("No Cookies were returned for refresh token")
				}
				testToken(t, cookies[0].Value, helpers.AdminUid)
			}
		})
	}
}

func TestRefreshTokenIpRange(t *testing.T) {

	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/auth/refresh",
	}

	tokenService := jwt_service.NewTokenService()

	// Create enterprise with Ip ranges set
	enterpriseWithIpRange := &models.Enterprise{
		UserIpRanges: "127.0.0.0/30",
	}
	factories.NewEnterpriseWithDefaults(ts.S.Db, enterpriseWithIpRange)

	userWithRange := &models.User{
		EnterpriseUID: enterpriseWithIpRange.Uid,
	}
	factories.NewUser(ts.S.Db, userWithRange)

	withRangeToken, withRangeExpiry, err := tokenService.CreateUserRefreshToken(userWithRange, userWithRange.EnterpriseUID, ts.Enterprise.RefreshTokenDuration)
	if err != nil {
		t.Fatalf("Error creating refresh token: %v", err)
	}
	withRangeCookie := tokenService.GetRefreshCookie(withRangeToken, *withRangeExpiry)

	cases := []helpers.TestCase{
		{
			TestName: "Can refresh with valid Ip address",
			Request:  request,
			RequestHeaders: map[string]string{
				"X-Forwarded-For": "127.0.0.1",
			},
			RequestCookies: []*http.Cookie{withRangeCookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyParts: []string{
					fmt.Sprint(enterpriseWithIpRange.EnterpriseDisplayName), // The enterprise Name
					userWithRange.Name,
					userWithRange.Uid,
					`"mdm_version"`,
				},
			},
		},
		{
			TestName: "Cannot refresh with invalid Ip address",
			Request:  request,
			RequestHeaders: map[string]string{
				"X-Forwarded-For": "127.0.0.200",
			},
			RequestCookies: []*http.Cookie{withRangeCookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyParts: []string{
					"Ip address is not allowed by enterprise",
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			res := ts.ExecuteTestCase(&test)

			// Validate the results from the test case
			ValidateResults(t, test, res)

			if res.Code == http.StatusOK {
				var authResponse responses.RefreshResponse
				_ = json.Unmarshal([]byte(res.Body.String()), &authResponse)

				testToken(t, authResponse.AccessToken, userWithRange.Uid)

				// Test the refresh token in the cookie
				cookies := res.Result().Cookies()
				if len(cookies) == 0 {
					t.Fatalf("No Cookies were returned for refresh token")
				}
				testToken(t, cookies[0].Value, userWithRange.Uid)
			}
		})
	}
}

func TestLogout(t *testing.T) {

	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/auth/logout",
	}

	tokenService := jwt_service.NewTokenService()

	// Generate good token for Admin
	adminRefreshCookie := tokenService.GetRefreshCookie(ts.AdminRefreshToken, time.Now().Add(1*time.Hour))

	cases := []helpers.TestCase{
		{
			TestName:       "Logout success",
			Request:        request,
			RequestCookies: []*http.Cookie{adminRefreshCookie},
			Expected: helpers.ExpectedResponse{
				StatusCode: 200,
				BodyParts: []string{
					"Logged Out", // The enterprise Name
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			res := ts.ExecuteTestCase(&test)

			ValidateResults(t, test, res)

			if res.Code == http.StatusOK {

				cookies := res.Result().Cookies()

				if len(cookies) == 0 {
					t.Fatalf("No Cookies were returned")
				}

				if cookies[0].Expires.After(time.Now()) {
					t.Fatalf("Cookie should set expiry to already expired")
				}
			}

		})
	}
}

func TestForgotPassword(t *testing.T) {
	ts.ClearTable("password_resets")
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/forgot-password",
	}

	cases := []helpers.TestCase{
		{
			TestName:    "Password reset not created for missing user",
			Request:     request,
			RequestBody: requests.ForgotPasswordRequest{Email: "user.not.exists@test.com"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Password reset email sent",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Model: &models.PasswordReset{},
						Scope: func(db *gorm.DB) *gorm.DB {
							return db.Where("user_id = 0")
						},
						CountExpected: 0,
					},
				},
			},
		},
		{
			TestName:    "Handles invalid email address",
			Request:     request,
			RequestBody: requests.ForgotPasswordRequest{Email: "this-is-not-an-email-address"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Email must be a valid email address",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Model: &models.PasswordReset{},
						Scope: func(db *gorm.DB) *gorm.DB {
							return db.Where("user_id = 0")
						},
						CountExpected: 0,
					},
				},
			},
		},
		{
			TestName:    "Handles missing email address",
			Request:     request,
			RequestBody: requests.ForgotPasswordRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Email is a required field",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Model: &models.PasswordReset{},
						Scope: func(db *gorm.DB) *gorm.DB {
							return db.Where("user_id = 0")
						},
						CountExpected: 0,
					},
				},
			},
		},
		{
			TestName: "Handles enterprise with forgot password disabled",
			Request:  request,
			RequestBody: requests.ForgotPasswordRequest{
				Email: ts.AdminUser.Email,
			},
			Setup: func(testCase *helpers.TestCase) {
				ts.Enterprise.ForgotPasswordEnabled = false
				ts.S.Db.Save(ts.Enterprise)
			},
			Teardown: func() {
				ts.Enterprise.ForgotPasswordEnabled = true
				ts.S.Db.Save(ts.Enterprise)
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Password reset email sent",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Name:          "Password Reset is not created in DB",
						Model:         &models.PasswordReset{UserUid: ts.AdminUser.Uid},
						CountExpected: 0,
					},
				},
				ExpectedCallBack: func(_ *httptest.ResponseRecorder) {
					// Check email was sent. Emails are sent in a separate go routine, so need to wait for that
					ts.S.Dependencies.GetMailer().Wg.Wait()
					assert.Equal(t, 1, mailer.MockCounter, "Count of emails should be 1")
					assert.Contains(t, mailer.MockEmails[1].HTML, "Forgot your password", "Expected email to contain forgot your password")
					assert.Contains(t, mailer.MockEmails[1].HTML, "Password resets are disabled for your enterprise. Please contact an administrator.", "Expected email to contain disabled message")
					assert.Contains(t, mailer.MockEmails[1].Recipients, ts.AdminUser.Email, "Expected email to be sent to users email address")
				},
			},
		},
		{
			TestName:    "Password reset created for user",
			Request:     request,
			RequestBody: requests.ForgotPasswordRequest{Email: ts.AdminUser.Email},
			Setup: func(testCase *helpers.TestCase) {
				ts.ClearTable("password_resets")
				err := os.Setenv("COUNTRY", "uk")
				if err != nil {
					log.Fatalf("error setting the country: %v", err)
				}
				ts.S.Dependencies.SetMailer(mailer.NewMailer(mailer.NewMockEmailDriver()))
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Password reset email sent",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Name:          "Password Reset is created in DB",
						Model:         &models.PasswordReset{UserUid: ts.AdminUser.Uid},
						CountExpected: 1,
					},
				},
				ExpectedCallBack: func(_ *httptest.ResponseRecorder) {
					// Check email was sent. Emails are sent in a separate go routine, so need to wait for that
					ts.S.Dependencies.GetMailer().Wg.Wait()
					assert.Equal(t, 1, mailer.MockCounter, "Count of emails should be 1")
					assert.Contains(t, mailer.MockEmails[1].HTML, "Forgot your password", "Expected email to contain forgot your password")
					assert.Contains(t, mailer.MockEmails[1].HTML, "/auth/reset/uk", "Expected email to contain forgot password link")
					assert.Contains(t, mailer.MockEmails[1].Recipients, ts.AdminUser.Email, "Expected email to be sent to users email address")
				},
			},
		},
		{
			TestName:    "Password reset created for user in Australia",
			Request:     request,
			RequestBody: requests.ForgotPasswordRequest{Email: ts.AdminUser.Email},
			Setup: func(testCase *helpers.TestCase) {
				ts.ClearTable("password_resets")
				err := os.Setenv("COUNTRY", "aus")
				if err != nil {
					log.Fatalf("error setting the country: %v", err)
				}
				ts.S.Dependencies.SetMailer(mailer.NewMailer(mailer.NewMockEmailDriver()))
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Password reset email sent",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Name:          "Password Reset is created in DB",
						Model:         &models.PasswordReset{UserUid: ts.AdminUser.Uid},
						CountExpected: 1,
					},
				},
				ExpectedCallBack: func(_ *httptest.ResponseRecorder) {
					// Check email was sent. Emails are sent in a separate go routine, so need to wait for that
					ts.S.Dependencies.GetMailer().Wg.Wait()
					assert.Equal(t, 1, mailer.MockCounter, "Count of emails should be 1")
					assert.Contains(t, mailer.MockEmails[1].HTML, "Forgot your password", "Expected email to contain forgot your password")
					assert.Contains(t, mailer.MockEmails[1].HTML, "/auth/reset/aus", "Expected email to contain forgot password link")
					assert.Contains(t, mailer.MockEmails[1].Recipients, ts.AdminUser.Email, "Expected email to be sent to users email address")
				},
			},
		},
		{
			TestName:    "Existing Password reset for user don't get deleted",
			Request:     request,
			RequestBody: requests.ForgotPasswordRequest{Email: ts.TestUser.Email},
			Setup: func(testCase *helpers.TestCase) {
				ts.ClearTable("password_resets")

				existingReset := &models.PasswordReset{
					UserUid: ts.TestUser.Uid,
				}
				factories.NewPasswordReset(ts.S.Db, existingReset)
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Password reset email sent",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Model: &models.PasswordReset{UserUid: ts.TestUser.Uid},
						Scope: func(db *gorm.DB) *gorm.DB {
							return db.Unscoped().Where("deleted_at IS NULL")
						},
						CountExpected: 2,
					},
				},
				ExpectedCallBack: func(_ *httptest.ResponseRecorder) {
					// Check email was sent. Emails are sent in a separate go routine, so need to wait for that
					ts.S.Dependencies.GetMailer().Wg.Wait()
					assert.Equal(t, 1, mailer.MockCounter, "Count of emails should be 1")
					assert.Contains(t, mailer.MockEmails[1].HTML, "Forgot your password")
				},
			},
		},
		{
			TestName: "request is rejected if any of the fields exceeds the max amount of characters",
			Request:  request,
			RequestBody: requests.ForgotPasswordRequest{
				Email: fmt.Sprintf("a@%v.com", strings.Repeat("A", 195)),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Email must be a maximum of 200 characters in length",
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			mailer.Reset()

			RunUnauthorisedTestCase(t, test)
		})
	}
}

func TestResetPassword(t *testing.T) {
	ts.ClearTable("password_resets")
	ts.ClearTable("failed_logins")
	ts.SetupDefaultUsers()

	validResets := factories.NewPasswordResets(ts.S.Db, &models.PasswordReset{
		UserUid: ts.TestUser.Uid,
	}, 4)
	usedReset := factories.NewPasswordResets(ts.S.Db, &models.PasswordReset{
		UserUid: ts.TestUser.Uid,
		Used:    true,
	}, 1)[0]
	expiredReset := factories.NewPasswordResets(ts.S.Db, &models.PasswordReset{
		UserUid:    ts.TestUser.Uid,
		ValidUntil: time.Now().Add(time.Hour * -1),
	}, 1)[0]
	adminReset := factories.NewPasswordResets(ts.S.Db, &models.PasswordReset{
		UserUid: ts.AdminUser.Uid,
	}, 1)[0]

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/reset-password",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Password reset works with valid input",
			Request:  request,
			RequestBody: requests.PasswordResetRequest{
				Token:              adminReset.Token,
				NewPassword:        "NewPassw0rd$$",
				NewPasswordConfirm: "NewPassw0rd$$",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Password reset",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Model:         &models.PasswordReset{UserUid: adminReset.UserUid, Used: true},
						CountExpected: 1,
					},
					{
						Model: &models.User{Uid: ts.AdminUser.Uid},
						Scope: func(db *gorm.DB) *gorm.DB {
							return db.Where("password <> ?", ts.AdminUser.Password)
						},
						CountExpected: 1,
					},
				},
			},
		},
		{
			TestName: "Password reset doesn't work with used token",
			Request:  request,
			RequestBody: requests.PasswordResetRequest{
				Token:              usedReset.Token,
				NewPassword:        "NewPassw0rd$$",
				NewPasswordConfirm: "NewPassw0rd$$",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Token has been used",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Model: &models.User{ID: ts.TestUser.ID},
						Scope: func(db *gorm.DB) *gorm.DB {
							return db.Where("password = ?", ts.TestUser.Password)
						},
						CountExpected: 1,
					},
				},
			},
		},
		{
			TestName: "Password reset doesn't work with expired token",
			Request:  request,
			RequestBody: requests.PasswordResetRequest{
				Token:              expiredReset.Token,
				NewPassword:        "NewPassw0rd$$",
				NewPasswordConfirm: "NewPassw0rd$$",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Token has expired",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Model: &models.User{ID: ts.TestUser.ID},
						Scope: func(db *gorm.DB) *gorm.DB {
							return db.Where("password = ?", ts.TestUser.Password)
						},
						CountExpected: 1,
					},
				},
			},
		},
		{
			TestName: "Password reset doesn't work if new passwords don't match",
			Request:  request,
			RequestBody: requests.PasswordResetRequest{
				Token:              expiredReset.Token,
				NewPassword:        "NewPassw0rd$$",
				NewPasswordConfirm: "NewPassw0rd$____!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Confirmation password does not match",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Model: &models.User{ID: ts.TestUser.ID},
						Scope: func(db *gorm.DB) *gorm.DB {
							return db.Where("password = ?", ts.TestUser.Password)
						},
						CountExpected: 1,
					},
				},
			},
		},
		{
			TestName: "Password reset doesn't work with short password",
			Request:  request,
			RequestBody: requests.PasswordResetRequest{
				Token:              validResets[1].Token,
				NewPassword:        "NewPass",
				NewPasswordConfirm: "NewPass",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Password must be at least 12 characters long",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Model:         &models.PasswordReset{ID: validResets[1].ID, Used: true},
						CountExpected: 0,
					},
					{
						Model: &models.User{ID: ts.TestUser.ID},
						Scope: func(db *gorm.DB) *gorm.DB {
							return db.Where("password = ?", ts.TestUser.Password)
						},
						CountExpected: 1,
					},
				},
			},
		},
		{
			TestName: "Password reset attempt fails with 3 previous reset attempts",
			Request:  request,
			RequestBody: requests.PasswordResetRequest{
				Token:              validResets[2].Token,
				NewPassword:        "NewPassword_!123",
				NewPasswordConfirm: "NewPassword_!123",
			},
			Setup: func(t *helpers.TestCase) {
				// 3 failed logins, within 15 mins, from same IP
				factories.NewFailedLogins(ts.S.Db, &models.FailedLogin{
					IpAddress: "127.0.0.0",
					Email:     "reset-password",
				}, 3)
			},
			Teardown: func() {
				ts.ClearTable("failed_logins")
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Too many attempts, you have been blocked temporarily",
			},
		},
		{
			TestName: "Password reset attempt success if 3 previous attempts, but 1 different ip",
			Request:  request,
			RequestBody: requests.PasswordResetRequest{
				Token:              validResets[2].Token,
				NewPassword:        "NewPassword_!123",
				NewPasswordConfirm: "NewPassword_!123",
			},
			Setup: func(t *helpers.TestCase) {
				// 3 failed logins, within 15 mins, 2 from same IP
				factories.NewFailedLogins(ts.S.Db, &models.FailedLogin{
					IpAddress: "127.0.0.0",
					Email:     "reset-password",
				}, 2)
				factories.NewFailedLogins(ts.S.Db, &models.FailedLogin{
					IpAddress: "127.0.0.1",
					Email:     "reset-password",
				}, 1)
			},
			Teardown: func() {
				ts.ClearTable("failed_logins")
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Password reset",
			},
		},
		{
			TestName: "Password reset attempt success if 3 previous attempts, but 1 over 15 mins ago",
			Request:  request,
			RequestBody: requests.PasswordResetRequest{
				Token:              validResets[3].Token,
				NewPassword:        "NewPassword_!123",
				NewPasswordConfirm: "NewPassword_!123",
			},
			Setup: func(t *helpers.TestCase) {
				// 3 failed logins, within 15 mins, 2 from same IP
				factories.NewFailedLogins(ts.S.Db, &models.FailedLogin{
					IpAddress: "127.0.0.0",
					Email:     "reset-password",
				}, 2)
				factories.NewFailedLogins(ts.S.Db, &models.FailedLogin{
					IpAddress: "127.0.0.0",
					Email:     "reset-password",
					CreatedAt: time.Now().Add(-16 * time.Minute),
				}, 1)
			},
			Teardown: func() {
				ts.ClearTable("failed_logins")
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Password reset",
			},
		},
		{
			TestName: "request is rejected if any of the fields exceeds the max amount of characters",
			Request:  request,
			RequestBody: requests.PasswordResetRequest{
				Token:              string(make([]byte, 101)),
				NewPassword:        fmt.Sprintf("1!a%v", strings.Repeat("A", 470)),
				NewPasswordConfirm: string(make([]byte, 473)),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Token must be a maximum of 100 characters in length",
					"NewPassword must be a maximum of 472 characters in length",
					"NewPasswordConfirm must be a maximum of 472 characters in length",
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunUnauthorisedTestCase(t, test)
		})
	}
}

// testToken will return the UID from the user table, not the id. The UID can be used to look the user up.
func testToken(t *testing.T, tokenToParse, uidToMatch string) {
	t.Helper()

	token, _ := jwt.Parse(tokenToParse, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		var hmacSampleSecret []byte
		return hmacSampleSecret, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	tokenUid := claims["id"].(string)

	assert.Equal(t, uidToMatch, tokenUid)

}
