package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/server/requests"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAllClasses(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get list of classes (populated table)",
			Setup:    ts.SetupDefaultClasses,
			Request:  helpers.Request{Method: http.MethodGet, URL: "/classes"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Barbarian"`,
					`"name":"Ranger"`,
					`"name":"Wizard"`,
				},
			},
		},
		{
			TestName: "can get list of classes (empty table)",
			Setup: func() {
				ts.ClearTable("characters")
				ts.ClearTable("classes")
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/classes"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "[]",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestGetClass(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get class by id",
			Setup:    ts.SetupDefaultClasses,
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/classes/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Barbarian"`},
			},
		},
		{
			TestName: "get /classes/:id returns 404 not found on class id not in database",
			Setup:    ts.SetupDefaultClasses,
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/classes/100",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "record not found",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestGetClassFeatures(t *testing.T) {
	ts.ClearTable("class_features")
	ts.ClearTable("classes")
	ts.ClearTable("features")

	classOne := &models.Class{}
	factories.NewClass(ts.S.Db, classOne)
	featureOne := &models.Feature{}
	factories.NewFeature(ts.S.Db, featureOne)
	classFeatureOne := &models.ClassFeature{ClassID: classOne.ID, FeatureID: featureOne.ID}
	factories.NewClassFeature(ts.S.Db, classFeatureOne)

	noFeatures := &models.Class{}
	factories.NewClass(ts.S.Db, noFeatures)

	cases := []helpers.TestCase{
		{
			TestName: "Can get features for class",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/classes/%v/features", classOne.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   fmt.Sprintf(`"name":"%v"`, featureOne.Name),
			},
		},
		{
			TestName: "Can get empty response for class with no features",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/classes/%v/features", noFeatures.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "[]",
			},
		},
		{
			TestName: "Can get 404 response for invalid class id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/classes/invalid-id/features",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestUpdateClass(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can update class by id with valid json in request body",
			Setup:    ts.SetupDefaultClasses,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/classes/1",
			},
			RequestBody: requests.ClassRequest{
				Name:             "Test Name",
				ShortDescription: "Test Description",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Test Name"`, `"short_description":"Test Description"`},
			},
		},
		{
			TestName: "put /classes/:id returns 404 not found on class id not in database",
			Setup:    ts.SetupDefaultClasses,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/classes/100",
			},
			RequestBody: requests.ClassRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "record not found",
			},
		},
		{
			TestName: "put /classes/:id returns 400 bad request on empty request body",
			Setup:    ts.SetupDefaultClasses,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/classes/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid class request body",
			},
		},
		{
			TestName: "put /classes/:id no update on empty class in request body",
			Setup:    ts.SetupDefaultClasses,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/classes/1",
			},
			RequestBody: requests.ClassRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Barbarian"`},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
