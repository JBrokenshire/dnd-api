package helpers

import (
	"dnd-api/api"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"net/http/httptest"
)

type TestCase struct {
	Name               string
	Request            Request
	RequestBody        interface{}
	RequestReader      io.Reader
	RequestContentType string
	RequestCookies     []*http.Cookie
	RequestHeaders     map[string]string
	HandlerFunc        func(s *api.Server, c echo.Context) error
	QueryMock          *QueryMock
	Expected           ExpectedResponse
	AccessToken        string
	Setup              func(test *TestCase)
	Teardown           func()
}

type PathParam struct {
	Name  string
	Value string
}

type Request struct {
	Method    string
	URL       string
	PathParam *PathParam
}

type DatabaseCheck struct {
	Name          string
	Model         interface{}
	Scope         func(db *gorm.DB) *gorm.DB
	CountExpected int
	DebugQuery    bool
}

type MockReply []map[string]interface{}

type QueryMock struct {
	Query string
	Reply MockReply
}

type ExpectedResponse struct {
	StatusCode       int
	BodyPart         string
	BodyParts        []string
	BodyPartMissing  string
	BodyPartsMissing []string
	DatabaseCheck    *DatabaseCheck
	DatabaseChecks   []*DatabaseCheck
	ExpectedCallback func(res *httptest.ResponseRecorder)
}
