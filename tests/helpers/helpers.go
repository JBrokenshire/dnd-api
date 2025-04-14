package helpers

import (
	"dnd-api/db/models"
	"dnd-api/pkg/try"
	"dnd-api/services/jwt_service"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"strings"
	"testing"
)

// Declare global helper Variables

var AdminUser *models.User
var AdminToken string
var TestUser *models.User
var TestToken string

func ClearTable(db *gorm.DB, tableName string) {
	err := db.Exec(fmt.Sprintf("DELETE FROM %v", tableName)).Error
	if err != nil {
		log.Fatalf("You can't clear that table. Err: %v", err)
	}
	err = db.Exec(fmt.Sprintf("ALTER TABLE %v AUTO_INCREMENT = 1", tableName)).Error
	if err != nil {
		log.Fatalf("Error setting autoincrement. Err: %v", err)
	}
}

func SetupDefaultUsers(db *gorm.DB) {
	ClearTable(db, "users")
	// Create an admin user so that we have some basic credentials in the database
	CreateAdminUser(db)
	CreateTestUser(db)
}

// CreateAdminUser creates an admin user and an auth token.
func CreateAdminUser(db *gorm.DB) {

	user := &models.User{
		ID:       1,
		Username: AdminUsername,
		Password: DefaultPasswordHash, // Abcd1234$
		Admin:    true,
	}
	// Save in the user table () We have ot omit the roles here so it doesn't upset the data into the table.
	err := db.Create(user).Error
	if err != nil {
		log.Fatalf("Unable to creatre admin user: %v", err)
	}

	// Create a usable token and refresh token. NOTE: authenticated set to True automatically bypassing 2FA
	tokenService := jwt_service.TokenService{}
	access, _, _ := tokenService.CreateUserAccessToken(user, true)
	AdminToken = access
	//refreshToken, _, _ := tokenService.CreateUserRefreshToken(user, user.EnterpriseUID, TestEnterprise.RefreshTokenDuration)
	//AdminRefreshToken = refreshToken
}

// CreateTestUser creates an test user and an auth token.
func CreateTestUser(db *gorm.DB) {
	user := &models.User{
		ID:       2,
		Username: TestUsername,
		Password: DefaultPasswordHash, // Abcd1234$
	}
	// Save in the user table () We have ot omit the roles here so it doesn't upset the data into the table.
	db.Create(user)
	TestUser = user

	// Create a usable token.
	tokenService := jwt_service.TokenService{}
	access, _, _ := tokenService.CreateUserAccessToken(user, true)
	TestToken = access
}

func RunNoAuthenticationAndPermissionTests(t *testing.T, e *echo.Echo, method string, url string) {
	RunNotAuthenticatedTest(t, e, method, url)
	RunNoPermissionTest(t, e, method, url)
}

func RunNotAuthenticatedTest(t *testing.T, e *echo.Echo, method string, url string) {

	testName := fmt.Sprintf("Cannot access %v if unauthenticated", url)
	t.Run(testName, func(t *testing.T) {
		req, _ := http.NewRequest(method, url, nil)
		res := try.ExecuteRequest(e, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Contains(t, res.Body.String(), "missing or malformed jwt")
	})
}

func RunNoPermissionTest(t *testing.T, e *echo.Echo, method string, url string) {

	testName := fmt.Sprintf("Cannot access %v if user doesnt have permission", strings.Replace(url, "/", "", -1))

	t.Run(testName, func(t *testing.T) {
		req, _ := http.NewRequest(method, url, nil)
		SetDefaultUserAgent(req)
		res := ExecuteTestUserRequest(e, req)
		assert.Equal(t, http.StatusForbidden, res.Code)
		assert.Contains(t, res.Body.String(), "Missing Required Permission")
	})
}

func ExecuteTestUserRequest(e *echo.Echo, req *http.Request) *try.HijackableResponseRecorder {
	if TestUser == nil {
		log.Fatalf("Unable to ExecuteTestUserRequest as TesUser is nil")
	}
	SetDefaultUserAgent(req)
	req.Header.Add("Authorization", "Bearer "+TestToken)
	return try.ExecuteRequest(e, req)
}

func SetDefaultUserAgent(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
}

// RunAuthorisedTestCase Run the Authorised Tests and run all assertions specified in the TestCase
func RunAuthorisedTestCase(t *testing.T, e *echo.Echo, test *try.TestCase, db *gorm.DB) {
	res := ExecuteAuthorisedTestCase(t, e, test)

	//log.Println(utils.PrettyPrint(res))
	try.ValidateResults(t, test, res, db)
}

// ExecuteAuthorisedTestCase Makes a request, and returns the response from ExecuteAuthorisedRequest.
func ExecuteAuthorisedTestCase(t *testing.T, e *echo.Echo, testCase *try.TestCase) *try.HijackableResponseRecorder {
	// Perform any setup needed for the test case
	if testCase.Setup != nil {
		testCase.Setup(testCase)
	}
	req, err := try.GenerateRequest(testCase)
	if err != nil {
		t.Fatalf("Unable to generate request")
	}
	res := ExecuteAuthorisedRequest(t, e, req)
	// Perform any teardown needed for test case
	//if testCase.Teardown != nil {
	//	testCase.Teardown()
	//}
	return res
}

// ExecuteAuthorisedRequest makes an admin user if it doesn't already exist and then add it's auth token into the header for the request
func ExecuteAuthorisedRequest(t *testing.T, e *echo.Echo, req *http.Request) *try.HijackableResponseRecorder {
	if AdminUser == nil {
		t.Fatalf("Admin User is NIL. Can't run test")
	}
	req.Header.Add("Authorization", "Bearer "+AdminToken)
	return try.ExecuteRequest(e, req)
}
