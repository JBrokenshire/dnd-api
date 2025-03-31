package middleware_test

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	mw "purplevisits.com/mdm/pkg/middleware"
	"sync"
	"testing"
)

func TestRequestLoggerWithConfig(t *testing.T) {
	e := echo.New()

	wg := sync.WaitGroup{}

	loggerMiddleware := mw.CustomLogger()
	mw.CustomLogFunction = func(v middleware.RequestLoggerValues) {
		log.Printf("Custom log function was called! %v", v)
		assert.Equal(t, http.StatusTeapot, v.Status)
		assert.Equal(t, "/test", v.URI)
		assert.Equal(t, "GET", v.Method)
		wg.Done()
	}
	e.Use(loggerMiddleware)

	wg.Add(1)
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusTeapot, "OK")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	wg.Wait()
}
