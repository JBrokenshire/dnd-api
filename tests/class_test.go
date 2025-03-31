package tests

import (
	"dnd-api/api/requests"
	"dnd-api/db/factories"
	m "dnd-api/db/models"
	"dnd-api/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestClass_List(t *testing.T) {
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
			URL:    fmt.Sprintf("/classes%v", query),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:    "Can get classes",
			Request: getRequest(""),
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
			Name:    "Can get page 0 of classes",
			Request: getRequest("?page=0&page_size=1"),
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
			Name:    "Can get page 1 of classes",
			Request: getRequest("?page=1&page_size=1"),
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
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestClass_Get(t *testing.T) {
	ts.ClearTable("classes")

	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			URL:    fmt.Sprintf("/classes/%v", id),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:    "Can't get class that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't get class with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can get class",
			Request: getRequest(class.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestClass_Create(t *testing.T) {
	ts.ClearTable("classes")

	request := helpers.Request{
		Method: http.MethodPost,
		URL:    "/classes",
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't create class without required fields",
			Request:     request,
			RequestBody: requests.CreateClassRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Name is a required field",
				},
			},
		},
		{
			Name:    "Can't create class if fields exceed max length",
			Request: request,
			RequestBody: requests.CreateClassRequest{
				Name: string(make([]byte, 201)),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Name must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can create class",
			Request: request,
			RequestBody: requests.CreateClassRequest{
				Name: "Test Class",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					`"name":"Test Class"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Class was created",
					Model: m.Class{
						Name: "Test Class",
					},
					CountExpected: 1,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestClass_Update(t *testing.T) {
	ts.ClearTable("classes")

	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			URL:    fmt.Sprintf("/classes/%v", id),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't update class without required fields",
			Request:     getRequest(class.ID),
			RequestBody: requests.UpdateClassRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Name is a required field",
				},
			},
		},
		{
			Name:    "Can't update class if fields exceed max length",
			Request: getRequest(class.ID),
			RequestBody: requests.UpdateClassRequest{
				Name: string(make([]byte, 201)),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Name must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can't update class that doesn't exist",
			Request: getRequest(1000),
			RequestBody: requests.UpdateClassRequest{
				Name: "Updated Name",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't update class with invalid id",
			Request: getRequest("invalid-id"),
			RequestBody: requests.UpdateClassRequest{
				Name: "Updated Name",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can update class",
			Request: getRequest(class.ID),
			RequestBody: requests.UpdateClassRequest{
				Name: "Updated Name",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Updated Name"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Class was updated",
					Model: m.Class{
						Name: "Updated Name",
					},
					CountExpected: 1,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestClass_Delete(t *testing.T) {
	ts.ClearTable("classes")

	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodDelete,
			URL:    fmt.Sprintf("/classes/%v", id),
		}
	}

	cases := []helpers.TestCase{
		{
			Name:    "Can't delete class that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't delete class with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can delete class",
			Request: getRequest(class.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Class deleted successfully",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Class was deleted",
					Model: m.Class{
						Name: class.Name,
					},
					CountExpected: 0,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
