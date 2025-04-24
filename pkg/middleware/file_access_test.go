package middleware

import (
	"dnd-api/services/file_service"
	"github.com/stretchr/testify/assert"
	"net/http"
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
			Name:           "Cannot access path traversal",
			RequestedUrl:   "public/../enterprises/abc-123/files/test",
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
