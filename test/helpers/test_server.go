package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/JBrokenshire/dnd-api/api"
	"github.com/JBrokenshire/dnd-api/api/routes"
	"github.com/JBrokenshire/dnd-api/db/migrations/process"
	"github.com/JBrokenshire/dnd-api/db/repositories"
	"github.com/JBrokenshire/dnd-api/pkg/dependencies"
	"github.com/JBrokenshire/dnd-api/pkg/validation"
	"github.com/JBrokenshire/dnd-api/test/mocks"
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

	ts.S.Dependencies = dependencies.NewDependencyService(db)

	// Add a wait group in to the deps service. Handlers will use this waitgroup for test cases when work is fired to a
	// separate go routine.
	ts.S.Dependencies.CreateWg()

	ts.S.Dependencies.SetFileStore(mocks.NewFileStoreMock())

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

func (ts *TestServer) GenerateRequest(testCase *TestCase) *http.Request {
	reqJson, err := json.Marshal(testCase.RequestBody)
	if err != nil {
		log.Printf("There was an error marshalling the json: %v", err)
	}

	req, _ := http.NewRequest(testCase.Request.Method, testCase.Request.Url, bytes.NewBuffer(reqJson))

	return req
}
