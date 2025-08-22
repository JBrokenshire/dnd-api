package tests

import (
	"dnd-api/tests/helpers"
	"dnd-api/tests/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFile_Get(t *testing.T) {
	fileStoreMock := mocks.NewFileStoreMock()
	ts.S.Dependencies.SetFileStore(fileStoreMock)

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/files/test-filepath",
	}

	cases := []helpers.TestCase{
		{
			Name: "Can't get file if there is an error from the file store",
			Setup: func(test *helpers.TestCase) {
				fileStoreMock.Reset()
				fileStoreMock.GetError = errors.New("test error")
			},
			Request: request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "Internal Server Error",
			},
		},
		{
			Name: "Can get file from file store",
			Setup: func(test *helpers.TestCase) {
				fileStoreMock.Reset()
			},
			Request: request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.GetCalls))
					assert.Contains(t, fileStoreMock.GetCalls[0], "test-filepath")
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
