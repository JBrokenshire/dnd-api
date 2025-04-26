package tests

import (
	"dnd-api/api/requests"
	"dnd-api/db/factories"
	m "dnd-api/db/models"
	"dnd-api/tests/helpers"
	"dnd-api/tests/mocks"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClass_List(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Create classes
	class := &m.Class{Name: "a"}
	factories.NewClass(ts.S.Db, class)
	class2 := &m.Class{Name: "b"}
	factories.NewClass(ts.S.Db, class2)
	namedClass := &m.Class{Name: "Test Class"}
	factories.NewClass(ts.S.Db, namedClass)

	// Create images
	logo := &m.File{Model: m.FileModelClassLogo, ModelId: class.ID}
	factories.NewFile(ts.S.Db, logo)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/classes%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get classes",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, namedClass.Name),
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
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, namedClass.Name),
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
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					fmt.Sprintf(`"name":"%v"`, namedClass.Name),
				},
			},
		},
		{
			Name:    "Can filter classes by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedClass.Name),
					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					fmt.Sprintf(`"name":"%v"`, class2.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestClass_Get(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("subclasses")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Create classes
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)
	class2 := &m.Class{}
	factories.NewClass(ts.S.Db, class2)

	// Create images
	classLogo := &m.File{Model: m.FileModelClassLogo, ModelId: class.ID}
	factories.NewFile(ts.S.Db, classLogo)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/classes/%v", id),
		}
	}

	permissionRequest := getRequest(class.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

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
					fmt.Sprintf(`"filename":"%v"`, classLogo.Filename),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class2.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestClass_Create(t *testing.T) {
	ts.ClearTable("classes")
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/classes",
	}

	RunNoAuthenticationTests(t, request.Method, request.Url)

	cases := []helpers.TestCase{
		{
			Name:        "Can't create class without required fields",
			Request:     request,
			RequestBody: requests.CreateClassRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name is a required field",
					"ShortDescription is a required field",
					"PrimaryAbility is a required field",
					"HitPointDieValue is a required field",
					"Saves is a required field",
				},
			},
		},
		{
			Name:    "Can't create class if the fields exceed max length",
			Request: request,
			RequestBody: requests.CreateClassRequest{
				Name:             string(make([]byte, 201)),
				PrimaryAbility:   string(make([]byte, 201)),
				Saves:            string(make([]byte, 201)),
				ShortDescription: "Test Short Description",
				HitPointDieValue: 6,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name must be a maximum of 200 characters in length",
					"PrimaryAbility must be a maximum of 200 characters in length",
					"Saves must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can create class",
			Request: request,
			RequestBody: requests.CreateClassRequest{
				Name:             "Test Class",
				ShortDescription: "Test Short Description",
				PrimaryAbility:   "Strength",
				HitPointDieValue: 6,
				Saves:            "Test Saves",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					`"name":"Test Class"`,
					`"short_description":"Test Short Description"`,
					`"primary_ability":"Strength"`,
					`"hit_point_die_value":6`,
					`"saves":"Test Saves"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Class was created",
					Model: m.Class{
						Name:             "Test Class",
						ShortDescription: "Test Short Description",
						PrimaryAbility:   "Strength",
						HitPointDieValue: 6,
						Saves:            "Test Saves",
					},
					CountExpected: 1,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestClass_Update(t *testing.T) {
	ts.ClearTable("classes")
	ts.SetupDefaultUsers()

	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			Url:    fmt.Sprintf("/classes/%v", id),
		}
	}

	permissionRequest := getRequest(class.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:        "Can't update class without required fields",
			Request:     getRequest(class.ID),
			RequestBody: requests.UpdateClassRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name is a required field",
					"ShortDescription is a required field",
					"PrimaryAbility is a required field",
					"HitPointDieValue is a required field",
					"Saves is a required field",
				},
			},
		},
		{
			Name:    "Can't update class if fields exceed max length",
			Request: getRequest(class.ID),
			RequestBody: requests.UpdateClassRequest{
				Name:             string(make([]byte, 201)),
				PrimaryAbility:   string(make([]byte, 201)),
				Saves:            string(make([]byte, 201)),
				ShortDescription: "Test Short Description",
				HitPointDieValue: 6,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name must be a maximum of 200 characters in length",
					"PrimaryAbility must be a maximum of 200 characters in length",
					"Saves must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can't update class that doesn't exist",
			Request: getRequest(1000),
			RequestBody: requests.UpdateClassRequest{
				Name:             "Test Class",
				ShortDescription: "Test Short Description",
				PrimaryAbility:   "Strength",
				HitPointDieValue: 6,
				Saves:            "Strength & Constitution",
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
				Name:             "Test Class",
				ShortDescription: "Test Short Description",
				PrimaryAbility:   "Strength",
				HitPointDieValue: 6,
				Saves:            "Strength & Constitution",
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
				Name:             "Test Class",
				ShortDescription: "Test Short Description",
				PrimaryAbility:   "Strength",
				HitPointDieValue: 6,
				Saves:            "Test Saves",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Test Class"`,
					`"short_description":"Test Short Description"`,
					`"primary_ability":"Strength"`,
					`"hit_point_die_value":6`,
					`"saves":"Test Saves"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Class was created",
					Model: m.Class{
						Name:             "Test Class",
						ShortDescription: "Test Short Description",
						PrimaryAbility:   "Strength",
						HitPointDieValue: 6,
						Saves:            "Test Saves",
					},
					CountExpected: 1,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestClass_Delete(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Set mocks
	fileStoreMock := mocks.NewFileStoreMock()
	ts.S.Dependencies.SetFileStore(fileStoreMock)

	setup := func(test *helpers.TestCase) {
		fileStoreMock.Reset()
	}

	// Create classes
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create images
	logo := &m.File{Model: m.FileModelClassLogo, ModelId: class.ID}
	factories.NewFile(ts.S.Db, logo)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodDelete,
			Url:    fmt.Sprintf("/classes/%v", id),
		}
	}

	permissionRequest := getRequest(class.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

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
			Setup:   setup,
			Request: getRequest(class.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					"Class deleted successfully",
				},
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Name: "Class was deleted",
						Model: m.Class{
							ID:   class.ID,
							Name: class.Name,
						},
						CountExpected: 0,
					},
					{
						Name: "Logo was deleted",
						Model: m.File{
							ID:      logo.ID,
							Model:   m.FileModelClassLogo,
							ModelId: class.ID,
						},
					},
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.DeleteCalls))
					assert.Equal(t, fmt.Sprintf("%v/%v", logo.FileLocation, logo.Filename), fileStoreMock.DeleteCalls[0])
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestClass_UploadLogo(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Set mocks
	fileStoreMock := mocks.NewFileStoreMock()
	ts.S.Dependencies.SetFileStore(fileStoreMock)

	setup := func(test *helpers.TestCase) {
		ts.ClearTable("files")
		fileStoreMock.Reset()
	}

	// Create class
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// PNG file
	pngBody, pngMw := createMultipartFile(t, "file", "../assets/example.png")
	// JPG file
	jpgBody, jpgMw := createMultipartFile(t, "file", "../assets/example.jpg")
	// JPEG file
	jpegBody, jpegMw := createMultipartFile(t, "file", "../assets/example.jpeg")
	// WEBP file
	webpBody, webpMw := createMultipartFile(t, "file", "../assets/example.webp")
	// PDF file
	pdfBody, pdfMw := createMultipartFile(t, "file", "../assets/example.pdf")

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPost,
			Url:    fmt.Sprintf("/classes/%v/upload/logo", id),
		}
	}

	permissionRequest := getRequest(class.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't upload image for class that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't upload image for class with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't upload image for class if no file is provided",
			Request: getRequest(class.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Unable to read file",
			},
		},
		{
			Name:               "Can't upload image for class with invalid file type",
			Request:            getRequest(class.ID),
			RequestReader:      pdfBody,
			RequestContentType: pdfMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Invalid file type",
			},
		},
		{
			Name:               "Can upload png",
			Setup:              setup,
			Request:            getRequest(class.ID),
			RequestReader:      pngBody,
			RequestContentType: pngMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelClassLogo,
						ModelId: class.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("classes/%v", class.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".png")
				},
			},
		},
		{
			Name:               "Can upload jpg",
			Setup:              setup,
			Request:            getRequest(class.ID),
			RequestReader:      jpgBody,
			RequestContentType: jpgMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelClassLogo,
						ModelId: class.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("classes/%v", class.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".jpg")
				},
			},
		},
		{
			Name:               "Can upload jpeg",
			Setup:              setup,
			Request:            getRequest(class.ID),
			RequestReader:      jpegBody,
			RequestContentType: jpegMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelClassLogo,
						ModelId: class.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("classes/%v", class.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".jpeg")
				},
			},
		},
		{
			Name:               "Can upload webp",
			Setup:              setup,
			Request:            getRequest(class.ID),
			RequestReader:      webpBody,
			RequestContentType: webpMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelClassLogo,
						ModelId: class.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("classes/%v", class.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".webp")
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}
