package middleware

import "github.com/labstack/echo/v4"

// FilesCache adds cache values for FILES
func FilesCache() echo.MiddlewareFunc {
	var filesCacheHeaders = map[string]string{
		"Cache-Control": "max-age=31536000, immutable",
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// Set our Cache headers
			res := c.Response()
			for k, v := range filesCacheHeaders {
				res.Header().Set(k, v)
			}
			return next(c)
		}
	}
}
