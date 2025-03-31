package jwt_service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateExtractors(t *testing.T) {
	e := echo.New()

	tests := []struct {
		name      string
		lookups   string
		setupFunc func(req *http.Request, c *echo.Context)
		expected  []string
		err       error
	}{
		{
			name:    "Header extraction",
			lookups: "header:Authorization:Bearer ",
			setupFunc: func(req *http.Request, c *echo.Context) {
				req.Header.Set(echo.HeaderAuthorization, "Bearer token123")
			},
			expected: []string{"token123"},
			err:      nil,
		},
		{
			name:    "Query extraction",
			lookups: "query:token",
			setupFunc: func(req *http.Request, c *echo.Context) {
				q := req.URL.Query()
				q.Add("token", "queryToken")
				req.URL.RawQuery = q.Encode()
			},
			expected: []string{"queryToken"},
			err:      nil,
		},
		{
			name:    "Param extraction",
			lookups: "param:id",
			setupFunc: func(req *http.Request, c *echo.Context) {
				e.GET("/:id", func(c echo.Context) error {
					return nil
				})
				req = httptest.NewRequest(http.MethodGet, "/123", nil)
				rec := httptest.NewRecorder()
				newContext := e.NewContext(req, rec)
				newContext.SetParamNames("id")
				newContext.SetParamValues("123")
				*c = newContext
			},
			expected: []string{"123"},
			err:      nil,
		},
		{
			name:    "Cookie extraction",
			lookups: "cookie:session",
			setupFunc: func(req *http.Request, c *echo.Context) {
				req.AddCookie(&http.Cookie{Name: "session", Value: "cookieValue"})
			},
			expected: []string{"cookieValue"},
			err:      nil,
		},
		{
			name:    "Form extraction",
			lookups: "form:username",
			setupFunc: func(req *http.Request, c *echo.Context) {
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
				err := req.ParseForm()
				if err != nil {
					t.Fatalf("Error parsing form: %v", err)
				}
				req.Form.Add("username", "formUser")
			},
			expected: []string{"formUser"},
			err:      nil,
		},
		{
			name:    "Header extraction missing value",
			lookups: "header:Authorization",
			setupFunc: func(req *http.Request, c *echo.Context) {
				// No header set
			},
			expected: nil,
			err:      errHeaderExtractorValueMissing,
		},
		{
			name:    "Query extraction missing value",
			lookups: "query:token",
			setupFunc: func(req *http.Request, c *echo.Context) {
				// No query set
			},
			expected: nil,
			err:      errQueryExtractorValueMissing,
		},
		{
			name:    "Param extraction missing value",
			lookups: "param:id",
			setupFunc: func(req *http.Request, c *echo.Context) {
				e.GET("/:id", func(c echo.Context) error {
					return nil
				})
				req = httptest.NewRequest(http.MethodGet, "/", nil)
				rec := httptest.NewRecorder()
				newContext := e.NewContext(req, rec)
				*c = newContext
			},
			expected: nil,
			err:      errParamExtractorValueMissing,
		},
		{
			name:    "Cookie extraction missing value",
			lookups: "cookie:session",
			setupFunc: func(req *http.Request, c *echo.Context) {
				// No cookie set
			},
			expected: nil,
			err:      errCookieExtractorValueMissing,
		},
		{
			name:    "Form extraction missing value",
			lookups: "form:username",
			setupFunc: func(req *http.Request, c *echo.Context) {
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
				err := req.ParseForm()
				if err != nil {
					t.Fatalf("Error parsing form: %v", err)
				}
			},
			expected: nil,
			err:      errFormExtractorValueMissing,
		},
		{
			name:    "Multiple extractors",
			lookups: "header:Authorization:Bearer,query:token",
			setupFunc: func(req *http.Request, c *echo.Context) {
				req.Header.Set(echo.HeaderAuthorization, "Bearer token123")
				q := req.URL.Query()
				q.Add("token", "queryToken")
				req.URL.RawQuery = q.Encode()
			},
			expected: []string{" token123", "queryToken"},
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if tt.setupFunc != nil {
				tt.setupFunc(req, &c)
			}

			extractors, err := CreateExtractors(tt.lookups)
			assert.NoError(t, err)

			var allValues []string
			for _, extractor := range extractors {
				values, err := extractor(c)
				assert.Equal(t, tt.err, err)
				allValues = append(allValues, values...)
			}

			assert.Equal(t, tt.expected, allValues)

		})
	}
}
