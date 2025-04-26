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

func TestRace_List(t *testing.T) {
	ts.ClearTable("races")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Create races
	race := &m.Race{Name: "a"}
	factories.NewRace(ts.S.Db, race)
	race2 := &m.Race{Name: "b"}
	factories.NewRace(ts.S.Db, race2)
	namedRace := &m.Race{Name: "Test Race"}
	factories.NewRace(ts.S.Db, namedRace)

	// Create images
	logo := &m.File{Model: m.FileModelRaceLogo, ModelId: race.ID}
	factories.NewFile(ts.S.Db, logo)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/races%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get races",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, namedRace.Name),
					`"total_count":3`,
				},
			},
		},
		{
			Name:    "Can get page 0 of races",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, namedRace.Name),
				},
			},
		},
		{
			Name:    "Can get page 1 of races",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					fmt.Sprintf(`"name":"%v"`, namedRace.Name),
				},
			},
		},
		{
			Name:    "Can filter races by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedRace.Name),
					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
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

func TestRace_Get(t *testing.T) {
	ts.ClearTable("races")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Create races
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)
	race2 := &m.Race{}
	factories.NewRace(ts.S.Db, race2)

	// Create images
	logo := &m.File{Model: m.FileModelRaceLogo, ModelId: race.ID}
	factories.NewFile(ts.S.Db, logo)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/races/%v", id),
		}
	}

	permissionRequest := getRequest(race.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't get race that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't get race with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can get race",
			Request: getRequest(race.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"filename":"%v"`, logo.Filename),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, race2.Name),
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

func TestRace_Create(t *testing.T) {
	ts.ClearTable("races")
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/races",
	}

	RunNoAuthenticationTests(t, request.Method, request.Url)

	cases := []helpers.TestCase{
		{
			Name:        "Can't create race without required fields",
			Request:     request,
			RequestBody: requests.CreateRaceRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name is a required field",
					"ShortDescription is a required field",
					"CreatureType is a required field",
					"Size is a required field",
					"BaseSpeed is a required field",
				},
			},
		},
		{
			Name:    "Can't create race if the fields exceed max length",
			Request: request,
			RequestBody: requests.CreateRaceRequest{
				Name:             string(make([]byte, 201)),
				CreatureType:     string(make([]byte, 201)),
				Size:             string(make([]byte, 201)),
				ShortDescription: "Test Short Description",
				BaseSpeed:        30,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name must be a maximum of 200 characters in length",
					"CreatureType must be a maximum of 200 characters in length",
					"Size must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can create race",
			Request: request,
			RequestBody: requests.CreateRaceRequest{
				Name:             "Test Race",
				ShortDescription: "Test Short Description",
				CreatureType:     "Test Creature Type",
				Size:             "Test Size",
				BaseSpeed:        30,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					`"name":"Test Race"`,
					`"short_description":"Test Short Description"`,
					`"creature_type":"Test Creature Type"`,
					`"size":"Test Size"`,
					`"base_speed":30`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Race was created",
					Model: m.Race{
						Name:             "Test Race",
						ShortDescription: "Test Short Description",
						CreatureType:     "Test Creature Type",
						Size:             "Test Size",
						BaseSpeed:        30,
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

func TestRace_Update(t *testing.T) {
	ts.ClearTable("races")
	ts.SetupDefaultUsers()

	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			Url:    fmt.Sprintf("/races/%v", id),
		}
	}

	permissionRequest := getRequest(race.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:        "Can't update race without required fields",
			Request:     getRequest(race.ID),
			RequestBody: requests.UpdateRaceRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name is a required field",
					"ShortDescription is a required field",
					"CreatureType is a required field",
					"Size is a required field",
					"BaseSpeed is a required field",
				},
			},
		},
		{
			Name:    "Can't update race if fields exceed max length",
			Request: getRequest(race.ID),
			RequestBody: requests.CreateRaceRequest{
				Name:             string(make([]byte, 201)),
				CreatureType:     string(make([]byte, 201)),
				Size:             string(make([]byte, 201)),
				ShortDescription: "Test Short Description",
				BaseSpeed:        30,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name must be a maximum of 200 characters in length",
					"CreatureType must be a maximum of 200 characters in length",
					"Size must be a maximum of 200 characters in length",
				},
			},
		},
		{
			Name:    "Can't update race that doesn't exist",
			Request: getRequest(1000),
			RequestBody: requests.UpdateRaceRequest{
				Name:             "Test Race",
				ShortDescription: "Test Short Description",
				CreatureType:     "Test Creature Type",
				Size:             "Test Size",
				BaseSpeed:        30,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't update race with invalid id",
			Request: getRequest("invalid-id"),
			RequestBody: requests.UpdateRaceRequest{
				Name:             "Test Race",
				ShortDescription: "Test Short Description",
				CreatureType:     "Test Creature Type",
				Size:             "Test Size",
				BaseSpeed:        30,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can update race",
			Request: getRequest(race.ID),
			RequestBody: requests.UpdateRaceRequest{
				Name:             "Test Race",
				ShortDescription: "Test Short Description",
				CreatureType:     "Test Creature Type",
				Size:             "Test Size",
				BaseSpeed:        30,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Test Race"`,
					`"short_description":"Test Short Description"`,
					`"creature_type":"Test Creature Type"`,
					`"size":"Test Size"`,
					`"base_speed":30`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Race was updated",
					Model: m.Race{
						Name:             "Test Race",
						ShortDescription: "Test Short Description",
						CreatureType:     "Test Creature Type",
						Size:             "Test Size",
						BaseSpeed:        30,
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

func TestRace_Delete(t *testing.T) {
	ts.ClearTable("races")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Set mocks
	fileStoreMock := mocks.NewFileStoreMock()
	ts.S.Dependencies.SetFileStore(fileStoreMock)

	setup := func(test *helpers.TestCase) {
		fileStoreMock.Reset()
	}

	// Create races
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	// Create images
	logo := &m.File{Model: m.FileModelRaceLogo, ModelId: race.ID}
	factories.NewFile(ts.S.Db, logo)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodDelete,
			Url:    fmt.Sprintf("/races/%v", id),
		}
	}

	permissionRequest := getRequest(race.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't delete race that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't delete race with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can delete race",
			Setup:   setup,
			Request: getRequest(race.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					"Race deleted successfully",
				},
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Name: "Race was deleted",
						Model: m.Race{
							ID:   race.ID,
							Name: race.Name,
						},
						CountExpected: 0,
					},
					{
						Name: "Logo was deleted",
						Model: m.File{
							ID:      logo.ID,
							Model:   m.FileModelRaceLogo,
							ModelId: race.ID,
						},
						CountExpected: 0,
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

func TestRace_UploadLogo(t *testing.T) {
	ts.ClearTable("races")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Set mocks
	fileStoreMock := mocks.NewFileStoreMock()
	ts.S.Dependencies.SetFileStore(fileStoreMock)

	setup := func(test *helpers.TestCase) {
		ts.ClearTable("files")
		fileStoreMock.Reset()
	}

	// Create race
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

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
			Url:    fmt.Sprintf("/races/%v/upload/logo", id),
		}
	}

	permissionRequest := getRequest(race.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't upload image for race that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't upload image for race with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't upload image for race if no file is provided",
			Request: getRequest(race.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Unable to read file",
			},
		},
		{
			Name:               "Can't upload image for race with invalid file type",
			Request:            getRequest(race.ID),
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
			Request:            getRequest(race.ID),
			RequestReader:      pngBody,
			RequestContentType: pngMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelRaceLogo,
						ModelId: race.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("races/%v", race.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".png")
				},
			},
		},
		{
			Name:               "Can upload jpg",
			Setup:              setup,
			Request:            getRequest(race.ID),
			RequestReader:      jpgBody,
			RequestContentType: jpgMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelRaceLogo,
						ModelId: race.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("races/%v", race.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".jpg")
				},
			},
		},
		{
			Name:               "Can upload jpeg",
			Setup:              setup,
			Request:            getRequest(race.ID),
			RequestReader:      jpegBody,
			RequestContentType: jpegMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelRaceLogo,
						ModelId: race.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("races/%v", race.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".jpeg")
				},
			},
		},
		{
			Name:               "Can upload webp",
			Setup:              setup,
			Request:            getRequest(race.ID),
			RequestReader:      webpBody,
			RequestContentType: webpMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelRaceLogo,
						ModelId: race.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("races/%v", race.ID))
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
