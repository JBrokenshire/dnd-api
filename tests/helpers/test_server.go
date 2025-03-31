package helpers

import (
	"bytes"
	"dnd-api/api"
	"dnd-api/api/routes"
	"dnd-api/db/migrations/process"
	"dnd-api/db/repositories"
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
)

type TestServer struct {
	S *api.Server
}

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
