package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"purplevisits.com/mdm/pkg/file_service"
	"purplevisits.com/mdm/pkg/jwt_service"
	"purplevisits.com/mdm/api/responses"
	"strings"
)

func FileAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		owner := c.Get("user").(*jwt.Token)
		claims := owner.Claims.(*jwt_service.JwtFilesCustomClaims)
		paths := claims.Paths

		requestUri := c.Request().URL.Path
		requestUri = strings.TrimPrefix(requestUri, "/files/")
		requestedPath := requestUri

		userId := claims.ID
		c.Set("userId", userId)

		if hasAccess(c.Request().Method, requestedPath, paths) {
			return next(c)
		}

		// No access to the path provided so return an error
		log.Printf("WARNING: No access to file path requested. UserId: %v    Path: %v", claims.ID, requestedPath)
		return responses.MessageResponse(c, http.StatusNotFound, "Not Found")
	}
}

func hasAccess(method, requestedPath string, paths map[string]file_service.FilesAccess) bool {

	if strings.Contains(requestedPath, "..") {
		log.Printf("WARNING: path requested with potential path traversal.  Path: %v", requestedPath)
		return false
	}

	// Set the required access based on the request Method
	requiredAccess := file_service.FilesAccessRead
	if method == "POST" || method == "DELETE" {
		requiredAccess = file_service.FilesAccessReadWrite
	}

	// Check if they have access.
	for path, access := range paths {

		// If we're looking for RW and hte path is just R then return
		if requiredAccess == file_service.FilesAccessReadWrite && access == file_service.FilesAccessRead {
			continue
		}

		// If the path contains a * then needs to start with that path, otherwise it needs to match Exactly
		if strings.Contains(path, "*") {
			checkPath := strings.TrimSuffix(path, "*")
			if strings.HasPrefix(requestedPath, checkPath) {
				return true
			}
		} else {
			if path == requestedPath {
				return true
			}
		}
	}

	return false
}
