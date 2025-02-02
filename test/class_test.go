package tests

import (
	"fmt"
	"github.com/JBrokenshire/dnd-api/api/requests"
	"github.com/JBrokenshire/dnd-api/db/factories"
	m "github.com/JBrokenshire/dnd-api/db/models"
	"github.com/JBrokenshire/dnd-api/test/helpers"
	"net/http"
	"testing"
)

func TestClassList(t *testing.T) {
	ts.ClearTable("classes")

	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)
	class2 := &m.Class{}
	factories.NewClass(ts.S.Db, class2)
	class3 := &m.Class{}
	factories.NewClass(ts.S.Db, class3)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/classes%v", query),
		}
	}

	cases := []helpers.TestCase{
		{
			TestName: "Can list classes",
			Request:  getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, class3.Name),
					`"total_count":3`,
				},
			},
		},
		{
			TestName: "Can get page 0 of classes",
			Request:  getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, class3.Name),
				},
			},
		},
		{
			TestName: "Can get page 1 of classes",
			Request:  getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, class3.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestClassGet(t *testing.T) {
	ts.ClearTable("classes")

	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)
	class2 := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/classes/%v", id),
		}
	}

	cases := []helpers.TestCase{
		{
			TestName: "Can't get class that doesn't exist",
			Request:  getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Cannot find class",
			},
		},
		{
			TestName: "Can get class",
			Request:  getRequest(class.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"short_description":"%v"`, class.ShortDescription),
					fmt.Sprintf(`"long_description":"%v"`, class.LongDescription),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class2.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestClassCreate(t *testing.T) {
	ts.ClearTable("classes")

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/classes",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Can't create class with no request body",
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Required fields are missing or not valid:",
			},
		},
		{
			TestName: "Can create class",
			Request:  request,
			RequestBody: requests.ClassCreateRequest{
				Name:             "test class",
				ShortDescription: "test short description",
				LongDescription:  "test long description",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					`"name":"test class"`,
					`"short_description":"test short description"`,
					`"long_description":"test long description"`,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
