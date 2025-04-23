package file_service

import (
	"github.com/labstack/echo/v4"
)

type Service interface {
	ReadFileInfo(c echo.Context) (string, string, error)
	FileExtensionAllowed(fileName string, extensions []string) bool
	MIMETypeAllowed(mimeType string, allowedTypes []string) bool
}
