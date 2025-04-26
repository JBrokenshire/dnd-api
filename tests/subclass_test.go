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

func TestSubclass_List(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("subclasses")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Create classes
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)
	class2 := &m.Class{}
	factories.NewClass(ts.S.Db, class2)

	// Create subclasses
	subclass := &m.Subclass{Name: "a", ClassId: class.ID}
	factories.NewSubclass(ts.S.Db, subclass)
	subclass2 := &m.Subclass{Name: "b", ClassId: class.ID}
	factories.NewSubclass(ts.S.Db, subclass2)
	differentClassSubclass := &m.Subclass{ClassId: class2.ID}
	factories.NewSubclass(ts.S.Db, differentClassSubclass)
	namedSubclass := &m.Subclass{Name: "Test Subclass", ClassId: class.ID}
	factories.NewSubclass(ts.S.Db, namedSubclass)

	// Create images
	logo := &m.File{Model: m.FileModelSubclassLogo, ModelId: subclass.ID}
	factories.NewFile(ts.S.Db, logo)

	getRequest := func(id interface{}, query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/subclasses/%v%v", id, query),
		}
	}

	permissionRequest := getRequest(class.ID, "")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't get subclasses for class that doesn't exist",
			Request: getRequest(1000, ""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't get subclasses for class with invalid id",
			Request: getRequest("invalid-id", ""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can get subclasses for class",
			Request: getRequest(class.ID, ""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
					fmt.Sprintf(`"name":"%v"`, namedSubclass.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, differentClassSubclass.Name),
				},
			},
		},
		{
			Name:    "Can get page 0 of subclasses for class",
			Request: getRequest(class.ID, "?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, differentClassSubclass.Name),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
					fmt.Sprintf(`"name":"%v"`, namedSubclass.Name),
				},
			},
		},
		{
			Name:    "Can get page 1 of subclasses for class",
			Request: getRequest(class.ID, "?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, differentClassSubclass.Name),
					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					fmt.Sprintf(`"name":"%v"`, namedSubclass.Name),
				},
			},
		},
		{
			Name:    "Can search subclasses by name",
			Request: getRequest(class.ID, "?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedSubclass.Name),
					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, differentClassSubclass.Name),
					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
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

func TestSubclass_Create(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("subclasses")
	ts.SetupDefaultUsers()

	// Create classes
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/subclasses",
	}

	RunNoAuthenticationTests(t, request.Method, request.Url)

	cases := []helpers.TestCase{
		{
			Name:        "Can't create subclass without required fields",
			Request:     request,
			RequestBody: requests.CreateSubclassRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"ClassId is a required field",
					"Name is a required field",
					"ShortDescription is a required field",
				},
			},
		},
		{
			Name:    "Can't create subclass if fields exceed the max length",
			Request: request,
			RequestBody: requests.CreateSubclassRequest{
				ClassId:          class.ID,
				Name:             string(make([]byte, 201)),
				ShortDescription: "Test Short Description",
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
			Name:    "Can't create subclass for class that doesn't exist",
			Request: request,
			RequestBody: requests.CreateSubclassRequest{
				ClassId:          1000,
				Name:             "Test Name",
				ShortDescription: "Test Short Description",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can create subclass",
			Request: request,
			RequestBody: requests.CreateSubclassRequest{
				ClassId:          class.ID,
				Name:             "Test Name",
				ShortDescription: "Test Short Description",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					`"name":"Test Name"`,
					`"short_description":"Test Short Description"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Subclass was created",
					Model: m.Subclass{
						ClassId:          class.ID,
						Name:             "Test Name",
						ShortDescription: "Test Short Description",
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

func TestSubclass_UploadLogo(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("subclasses")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Set mocks
	fileStoreMock := mocks.NewFileStoreMock()
	ts.S.Dependencies.SetFileStore(fileStoreMock)

	setup := func(test *helpers.TestCase) {
		ts.ClearTable("files")
		fileStoreMock.Reset()
	}

	// Create classes
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)
	class2 := &m.Class{}
	factories.NewClass(ts.S.Db, class2)

	// Create subclasses
	subclass := &m.Subclass{ClassId: class.ID}
	factories.NewSubclass(ts.S.Db, subclass)
	differentClassSubclass := &m.Subclass{ClassId: class2.ID}
	factories.NewSubclass(ts.S.Db, differentClassSubclass)

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

	getRequest := func(classId, subclassId interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPost,
			Url:    fmt.Sprintf("/subclasses/%v/%v/upload/logo", classId, subclassId),
		}
	}

	permissionRequest := getRequest(class.ID, subclass.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't upload image for subclass with class id that doesn't exist",
			Request: getRequest(1000, subclass.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't upload image for subclass with invalid class id",
			Request: getRequest("invalid-id", subclass.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't upload image for subclass that doesn't exist",
			Request: getRequest(class.ID, 1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Subclass not found",
			},
		},
		{
			Name:    "Can't upload image for subclass with invalid id",
			Request: getRequest(class.ID, "invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Subclass not found",
			},
		},
		{
			Name:    "Can't upload image for subclass that belongs to a different class",
			Request: getRequest(class.ID, differentClassSubclass.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Subclass not found",
			},
		},
		{
			Name:    "Can't upload image for subclass if no file is provided",
			Request: getRequest(class.ID, subclass.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Unable to read file",
			},
		},
		{
			Name:               "Can't upload image for subclass with invalid file type",
			Request:            getRequest(class.ID, subclass.ID),
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
			Request:            getRequest(class.ID, subclass.ID),
			RequestReader:      pngBody,
			RequestContentType: pngMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelSubclassLogo,
						ModelId: subclass.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("classes/%v/subclasses/%v", class.ID, subclass.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".png")
				},
			},
		},
		{
			Name:               "Can upload jpg",
			Setup:              setup,
			Request:            getRequest(class.ID, subclass.ID),
			RequestReader:      jpgBody,
			RequestContentType: jpgMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelSubclassLogo,
						ModelId: subclass.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("classes/%v/subclasses/%v", class.ID, subclass.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".jpg")
				},
			},
		},
		{
			Name:               "Can upload jpeg",
			Setup:              setup,
			Request:            getRequest(class.ID, subclass.ID),
			RequestReader:      jpegBody,
			RequestContentType: jpegMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelSubclassLogo,
						ModelId: subclass.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("classes/%v/subclasses/%v", class.ID, subclass.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".jpeg")
				},
			},
		},
		{
			Name:               "Can upload webp",
			Setup:              setup,
			Request:            getRequest(class.ID, subclass.ID),
			RequestReader:      webpBody,
			RequestContentType: webpMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelSubclassLogo,
						ModelId: subclass.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("classes/%v/subclasses/%v", class.ID, subclass.ID))
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
