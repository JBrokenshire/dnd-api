package helpers

import (
	"bytes"
	"dnd-api/api"
	"dnd-api/api/routes"
	"dnd-api/db/migrations/process"
	m "dnd-api/db/models"
	"dnd-api/db/repositories"
	"dnd-api/pkg/jwt_service"
	"dnd-api/pkg/validation"
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
	"time"
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
	AdminEmail                  = "jbrokenshire0306@gmail.com"
	AdminPassword               = "Abcd1234$"
	AdminUid                    = "4758ad4a-73ea-4d91-b6bb-eca1fd12f015"
	DefaultPasswordHash         = "$2a$04$VownpTI87qYPnFKzXCy1vO.76A3LnHkADRlO/rm4bmuo9Ze7SkalW" // Abcd1234$
	TestPassword                = "Abcd1234$"
	TestClientSecretHash        = "$2a$04$mH3cYi8E4oiKXWz7QiZ63.g5lOlq0kH9C.inhtXGkZ8701SXk7xmW" // Nw8OP16TYSNkwua9bBXAiRczvLmStu5CVKxycssfrZLncCJF3aHGQKt-X4jjmgXXnVdDqwkm
	TestClientSecret            = "Nw8OP16TYSNkwua9bBXAiRczvLmStu5CVKxycssfrZLncCJF3aHGQKt-X4jjmgXXnVdDqwkm"
	RefreshTokenDurationMinutes = 10080 // 7 days
)

func NewTestServer(envFileLocation string) *TestServer {
	err := godotenv.Load(envFileLocation)
	if err != nil {
		log.Println("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_EXPOSE_DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Unable to connect to test database: %v\n", err.Error())
	}

	db.LogMode(false)

	ts := &TestServer{
		S: &api.Server{
			Echo:  echo.New(),
			Db:    db,
			Repos: repositories.NewRepos(db),
		},
	}

	ts.migrateDatabase()
	ts.S.Echo.Validator = validation.NewCustomValidator(validator.New())
	routes.ConfigureRoutes(ts.S)

	return ts
}

func (ts *TestServer) migrateDatabase() {
	_ = os.Setenv("DB_HOST", os.Getenv("TEST_DB_HOST"))
	_ = os.Setenv("DB_PORT", os.Getenv("TEST_EXPOSE_DB_PORT"))
	process.Run()
}

// ExecuteTestCase Makes a request, and returns the response from ExecuteRequest
func (ts *TestServer) ExecuteTestCase(test *TestCase) *httptest.ResponseRecorder {
	// Perform any setup needed for the test case
	if test.Setup != nil {
		test.Setup(test)
	}

	request := ts.GenerateRequest(test)
	response := ts.ExecuteRequest(request)

	// Perform any teardown needed for the test case
	if test.Teardown != nil {
		test.Teardown()
	}

	return response
}

// ExecuteRequest Executes a request against the API. This runs it locally against the handler
func (ts *TestServer) ExecuteRequest(request *http.Request) *httptest.ResponseRecorder {
	// Create a new recorder then process the request with the server
	rr := httptest.NewRecorder()
	ts.S.Echo.ServeHTTP(rr, request)
	return rr
}

// ClearTable Clear a table and reset the autoincrement
func (ts *TestServer) ClearTable(tableName string) {
	err := ts.S.Db.Exec(fmt.Sprintf("DELETE FROM %v", tableName)).Error
	if err != nil {
		log.Fatalf("Error clearing table %v: %v", tableName, err.Error())
	}
	err = ts.S.Db.Exec(fmt.Sprintf("ALTER TABLE %v AUTO_INCREMENT = 1", tableName)).Error
	if err != nil {
		log.Fatalf("Error clearing table %v: %v", tableName, err.Error())
	}
}

// SetupDefaultUsers Clear the users table and create the Admin and Test User
func (ts *TestServer) SetupDefaultUsers() {
	ts.ClearTable("users")
	ts.ClearTable("user_roles")
	// Create an admin user so that we have some basic credentials in the database
	ts.CreateAdminUser()
	ts.CreateTestUser()
}

// CreateAdminUser creates an admin user and an auth token.
func (ts *TestServer) CreateAdminUser() {
	user := &m.User{
		ID:         1,
		Uid:        AdminUid,
		Email:      AdminEmail,
		Name:       "Jared Brokenshire",
		Pronouns:   "He/Him",
		Password:   DefaultPasswordHash, // Abcd1234$
		SuperAdmin: true,
	}
	// Save in the user table () We have ot omit the roles here so it doesn't upsert the data into the table.
	ts.S.Db.Create(user)

	// Give this user the admin role
	ts.S.Db.Model(&m.User{ID: user.ID}).Association("Roles").Append([]m.Role{{ID: 1}})
	ts.AdminUser = user

	// Create a usable token and refresh token. NOTE: authenticated set to True automatically bypassing 2FA
	tokenService := jwt_service.TokenService{}
	access, _, _ := tokenService.CreateUserAccessToken(user, true)
	ts.AdminToken = access
	refreshToken, _, _ := tokenService.CreateUserRefreshToken(user, uint(10*time.Minute))
	ts.AdminRefreshToken = refreshToken

}

// CreateTestUser creates an test user and an auth token.
func (ts *TestServer) CreateTestUser() {
	user := &m.User{
		ID:       2,
		Uid:      "bdd72004-0786-4605-88b3-3d297c1f7b42",
		Email:    "testing@email.com",
		Password: DefaultPasswordHash, // Abcd1234$
		Name:     "Test User",
	}
	// Save in the user table () We have ot omit the roles here so it doesn't upsert the data into the table.
	ts.S.Db.Create(user)
	ts.TestUser = user

	// Create a usable token.
	tokenService := jwt_service.TokenService{}
	access, _, _ := tokenService.CreateUserAccessToken(user, true)
	ts.TestToken = access
}

func (ts *TestServer) SetDefaultTestHeaders(req *http.Request) {
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderXRealIP, "127.0.0.0")
}

func (ts *TestServer) SetDefaultUserAgent(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
}

func (ts *TestServer) GenerateRequest(test *TestCase) *http.Request {
	reqJson, err := json.Marshal(test.RequestBody)
	if err != nil {
		log.Printf("There was an error marshalling the json for the request body: %v", err.Error())
	}

	var req *http.Request
	if test.RequestReader != nil {
		req, err = http.NewRequest(test.Request.Method, test.Request.URL, test.RequestReader)
	} else {
		req, err = http.NewRequest(test.Request.Method, test.Request.URL, bytes.NewBuffer(reqJson))
	}

	// Set IP address and default context type to JSON
	ts.SetDefaultTestHeaders(req)

	// Change content type if one is set
	if test.RequestContentType != "" {
		req.Header.Set(echo.HeaderContentType, test.RequestContentType)
	}

	// Add in a default user agent
	ts.SetDefaultUserAgent(req)

	// Add cookies if present
	if len(test.RequestCookies) > 0 {
		for _, cookie := range test.RequestCookies {
			req.AddCookie(cookie)
		}
	}

	// Add additional headers if present
	if len(test.RequestHeaders) > 0 {
		for headerKey, headerValue := range test.RequestHeaders {
			// Set required to override content type. May need to be updated if you need multiple headers with the same key
			req.Header.Set(headerKey, headerValue)
		}
	}

	return req
}
