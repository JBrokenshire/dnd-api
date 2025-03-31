package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"purplevisits.com/mdm/pkg/file_service"
	"testing"
)

func TestUserHasAccess(t *testing.T) {
	paths := map[string]file_service.FilesAccess{}
	paths["public/*"] = file_service.FilesAccessRead
	paths["enterprises/abc-123/files/*"] = file_service.FilesAccessRead

	testCases := []struct {
		Name           string
		RequestedUrl   string
		Method         string
		ExpectedResult bool
	}{
		{
			Name:           "Can access public folder",
			RequestedUrl:   "public/test-folder/1",
			Method:         http.MethodGet,
			ExpectedResult: true,
		},
		{
			Name:           "Can access enterprise folder",
			RequestedUrl:   "enterprises/abc-123/files/test",
			Method:         http.MethodGet,
			ExpectedResult: true,
		},
		{
			Name:           "Cannot access path traversal",
			RequestedUrl:   "public/../enterprises/abc-123/files/test",
			Method:         http.MethodGet,
			ExpectedResult: false,
		},
		{
			Name:           "Cannot access another enterprises folder",
			RequestedUrl:   "enterprises/other-enterprise/files/test",
			Method:         http.MethodGet,
			ExpectedResult: false,
		},
		{
			Name:           "Cannot access enterprise root folder",
			RequestedUrl:   "enterprises/abc-123/file.txt",
			Method:         http.MethodGet,
			ExpectedResult: false,
		},
		{
			Name:           "Cannot access folder they dont have access to",
			RequestedUrl:   "enterprises/abc-123/jobs/123/audio.mp3",
			Method:         http.MethodGet,
			ExpectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Helper()
			result := hasAccess(testCase.Method, testCase.RequestedUrl, paths)
			assert.Equal(t, testCase.ExpectedResult, result, "Expected access result to match")
		})
	}
}

func TestOwnerHasAccess(t *testing.T) {
	paths := map[string]file_service.FilesAccess{}
	paths["public/*"] = file_service.FilesAccessRead
	paths["enterprises/abc-123/files/possessions/1"] = file_service.FilesAccessRead
	paths["enterprises/abc-123/files/possessions/2"] = file_service.FilesAccessRead

	testCases := []struct {
		Name           string
		RequestedUrl   string
		Method         string
		ExpectedResult bool
	}{
		{
			Name:           "Can access public folder",
			RequestedUrl:   "public/test-folder/1",
			Method:         http.MethodGet,
			ExpectedResult: true,
		},
		{
			Name:           "Can access possession they own",
			RequestedUrl:   "enterprises/abc-123/files/possessions/1",
			Method:         http.MethodGet,
			ExpectedResult: true,
		},
		{
			Name:           "Cannot access path traversal",
			RequestedUrl:   "public/../enterprises/abc-123/files/test",
			Method:         http.MethodGet,
			ExpectedResult: false,
		},
		{
			Name:           "Cannot access another enterprises folder",
			RequestedUrl:   "enterprises/other-enterprise/files/possession/1",
			Method:         http.MethodGet,
			ExpectedResult: false,
		},
		{
			Name:           "Cannot access possession they do not own",
			RequestedUrl:   "enterprises/abc-123/files/possessions/3/pic/png",
			Method:         http.MethodGet,
			ExpectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Helper()
			result := hasAccess(testCase.Method, testCase.RequestedUrl, paths)
			assert.Equal(t, testCase.ExpectedResult, result, "Expected access result to match")
		})
	}
}
