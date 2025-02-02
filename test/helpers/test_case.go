package helpers

import (
	"github.com/jinzhu/gorm"
	"net/http/httptest"
)

const UserId = 1

type TestCase struct {
	TestName    string
	Request     Request
	RequestBody interface{}
	Expected    ExpectedResponse
	Setup       func(testCase *TestCase)
	Teardown    func()
}

type Request struct {
	Method string
	Url    string
}

type DatabaseCheck struct {
	Name          string
	Model         interface{}
	Scope         func(*gorm.DB) *gorm.DB
	CountExpected int
	DebugQuery    bool
}

type ExpectedResponse struct {
	StatusCode       int
	BodyPart         string
	BodyParts        []string
	BodyPartMissing  string
	BodyPartsMissing []string
	DatabaseCheck    *DatabaseCheck
	DatabaseChecks   []*DatabaseCheck
	ExpectedCallBack func(res *httptest.ResponseRecorder)
}
