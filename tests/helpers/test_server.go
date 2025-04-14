package helpers

import (
	"bytes"
	"dnd-api/api"
	"dnd-api/api/routes"
	"dnd-api/config"
	"dnd-api/db/migrations/process"
	m "dnd-api/db/models"
	"dnd-api/db/repositories"
	"dnd-api/pkg/validation"
	"dnd-api/services/jwt_service"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type TestServer struct {
	S                 *api.Server
	AdminUser         *m.User
	AdminToken        string
	AdminRefreshToken string
	TestUser          *m.User
	TestToken         string
}

const (
	AdminUsername       = "JBrokenshire"
	AdminPassword       = "Abcd1234$"
	TestUsername        = "TestUser"
	DefaultPasswordHash = "$2a$04$VownpTI87qYPnFKzXCy1vO.76A3LnHkADRlO/rm4bmuo9Ze7SkalW" // Abcd1234$
)

func NewTestServer(envFileLoc string) *TestServer {
	err := godotenv.Load(envFileLoc)
	if err != nil {
		log.Println("Error loading .env file")
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_EXPOSE_DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Unable to connect to test database: %v\n", err.Error())
	}

	// disable default logging of errors, we were going to be creating some on purpose!
	db.LogMode(false)

	ts := &TestServer{
		S: &api.Server{
			Echo:  echo.New(),
			Db:    db,
			Repos: repositories.NewRepos(db),
		},
	}

	// Update the config
	conf := config.Get()
	conf.HashCost = 4
	conf.BruteForceLimit = 3
	conf.LogMiddleware = false
	conf.LoginRateLimit = 1000

	ts.migrateDatabase()
	ts.S.Echo.Validator = validation.NewCustomValidator(validator.New())
	routes.ConfigureRoutes(ts.S)

	return ts
}

func (ts *TestServer) migrateDatabase() {
	// This project requires DB_PORT and DB_HOST loaded from .env automatically. Let's make them match
	_ = os.Setenv("DB_HOST", os.Getenv("TEST_DB_HOST"))
	_ = os.Setenv("DB_PORT", os.Getenv("TEST_EXPOSE_DB_PORT"))
	process.Run()
}

// ExecuteAuthorisedTestCase Makes a request, and returns the response from ExecuteAuthorisedRequest.
func (ts *TestServer) ExecuteAuthorisedTestCase(testCase *TestCase) *httptest.ResponseRecorder {
	// Perform any setup needed for the test case
	if testCase.Setup != nil {
		testCase.Setup(testCase)
	}
	req := ts.GenerateRequest(testCase)
	res := ts.ExecuteAuthorisedRequest(req)
	// Perform any teardown needed for test case
	if testCase.Teardown != nil {
		testCase.Teardown()
	}
	return res
}

// ExecuteAuthorisedRequest makes an admin user if it doesn't already exist and then add it's auth token into the header for the request
func (ts *TestServer) ExecuteAuthorisedRequest(req *http.Request) *httptest.ResponseRecorder {
	if ts.AdminUser == nil {
		ts.CreateAdminUser()
	}
	if req.Header.Get("Authorization") == "" {
		req.Header.Add("Authorization", "Bearer "+ts.AdminToken)
	}
	return ts.ExecuteRequest(req)
}

// ExecuteTestCase Makes a request, and returns the response from ExecuteRequest.
func (ts *TestServer) ExecuteTestCase(testCase *TestCase) *httptest.ResponseRecorder {
	// Perform any setup needed for the test case
	if testCase.Setup != nil {
		testCase.Setup(testCase)
	}
	req := ts.GenerateRequest(testCase)
	res := ts.ExecuteRequest(req)
	// Perform any teardown needed for test case
	if testCase.Teardown != nil {
		testCase.Teardown()
	}
	return res
}

// ExecuteRequest Executes a request against the API. THis runs it locally against the handler
func (ts *TestServer) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {

	// Create a new recorder then process request with server.
	rr := httptest.NewRecorder()
	ts.S.Echo.ServeHTTP(rr, req)
	return rr
}

func (ts *TestServer) ExecuteTestUserRequest(req *http.Request) *httptest.ResponseRecorder {
	if ts.TestUser == nil {
		ts.CreateTestUser()
	}
	ts.SetDefaultUserAgent(req)
	req.Header.Add("Authorization", "Bearer "+ts.TestToken)
	return ts.ExecuteRequest(req)
}

// ClearTable Clear a table and reset the autoincrement
func (ts *TestServer) ClearTable(tableName string) {
	err := ts.S.Db.Exec(fmt.Sprintf("DELETE FROM %v", tableName)).Error
	if err != nil {
		log.Fatalf("You can't clear that table. Err: %v", err)
	}
	err = ts.S.Db.Exec(fmt.Sprintf("ALTER TABLE %v AUTO_INCREMENT = 1", tableName)).Error
	if err != nil {
		log.Fatalf("Error setting autoincrement. Err: %v", err)
	}
}

// SetupDefaultUsers Clear the users table and create the Admin and Test User
func (ts *TestServer) SetupDefaultUsers() {
	ts.ClearTable("users")
	// Create an admin user so that we have some basic credentials in the database
	ts.CreateAdminUser()
	ts.CreateTestUser()
}

// CreateAdminUser creates an admin user and an auth token.
func (ts *TestServer) CreateAdminUser() {
	user := &m.User{
		ID:       1,
		Username: AdminUsername,
		Password: DefaultPasswordHash, // Abcd1234$
		Admin:    true,
	}
	// Save in the user table () We have ot omit the roles here so it doesn't upset the data into the table.
	ts.S.Db.Create(user)
	ts.AdminUser = user

	// Create a usable token and refresh token. NOTE: authenticated set to True automatically bypassing 2FA
	tokenService := jwt_service.TokenService{}
	access, _, _ := tokenService.CreateUserAccessToken(user, true)
	ts.AdminToken = access
	refreshToken, _, _ := tokenService.CreateUserRefreshToken(user)
	ts.AdminRefreshToken = refreshToken

}

// CreateTestUser creates an test user and an auth token.
func (ts *TestServer) CreateTestUser() {
	user := &m.User{
		ID:       2,
		Username: TestUsername,
		Password: DefaultPasswordHash, // Abcd1234$
	}
	// Save in the user table () We have ot omit the roles here so it doesn't upset the data into the table.
	ts.S.Db.Create(user)
	ts.TestUser = user

	// Create a usable token.
	tokenService := jwt_service.TokenService{}
	access, _, _ := tokenService.CreateUserAccessToken(user, true)
	ts.TestToken = access
}

// GetDb Return the database from the server
func (ts *TestServer) GetDb() *gorm.DB {
	return ts.S.Db
}

func (ts *TestServer) CreateOrDie(o interface{}) {
	err := ts.S.Db.Create(o).Error
	if err != nil {
		log.Panicf("Error creating object as part of a test: %v", err)
	}
}

func (ts *TestServer) SetDefaultTestHeaders(req *http.Request) {
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderXRealIP, "127.0.0.0")
}
func (ts *TestServer) SetDefaultUserAgent(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
}
func (ts *TestServer) GenerateRequest(testCase *TestCase) *http.Request {
	reqJson, err := json.Marshal(testCase.RequestBody)
	if err != nil {
		log.Printf("There was an error marshalling the json: %v", err)
	}

	var req *http.Request
	if testCase.RequestReader != nil {
		req, err = http.NewRequest(testCase.Request.Method, testCase.Request.Url, testCase.RequestReader)
	} else {
		req, err = http.NewRequest(testCase.Request.Method, testCase.Request.Url, bytes.NewBuffer(reqJson))
	}

	// Set IP address and default context type to JSON
	ts.SetDefaultTestHeaders(req)

	// Change Content Type if one is set
	if testCase.RequestContentType != "" {
		req.Header.Set(echo.HeaderContentType, testCase.RequestContentType)
	}

	// Add in a default user agent.
	ts.SetDefaultUserAgent(req)

	// Add cookies in if present
	if len(testCase.RequestCookies) > 0 {
		for _, cookie := range testCase.RequestCookies {
			req.AddCookie(cookie)
		}
	}

	if len(testCase.RequestHeaders) > 0 {
		for headerKey, headerValue := range testCase.RequestHeaders {
			// Set required to override content type. May need to be updated if you need multiple headers
			// with the same key.
			req.Header.Set(headerKey, headerValue)
		}
	}

	return req
}
