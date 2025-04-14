package file_service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"slices"
	"strings"
)

type FilesAccess string

var (
	FilesAccessRead      FilesAccess = "R"
	FilesAccessReadWrite FilesAccess = "RW"
)

type FilesService struct {
	Db *gorm.DB
}

func NewFilesService(db *gorm.DB) *FilesService {
	return &FilesService{
		Db: db,
	}
}

func (s *FilesService) ReadFileInfo(c echo.Context) (string, string, error) {

	r := c.Request()
	w := c.Response()
	if err := r.ParseMultipartForm(10485); err != nil {
		return "", "", fmt.Errorf("error parsing multipart form: %v", err)
	}

	// Limit upload size
	r.Body = http.MaxBytesReader(w, r.Body, 10485)

	//
	file, multipartFileHeader, err := r.FormFile("file")

	if err != nil {
		return "", "", fmt.Errorf("error getting form file from request: %v", err)
	}

	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		return "", "", fmt.Errorf("error reading the file header: %v", err)
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		return "", "", fmt.Errorf("error seeking file: %v", err)
	}

	if multipartFileHeader == nil {
		return "", "", fmt.Errorf("multipartFileHeader is nil. Cant read file header")
	}

	return multipartFileHeader.Filename, http.DetectContentType(fileHeader), nil
}

func (s *FilesService) FileExtensionAllowed(fileName string, extensions []string) bool {
	for _, extension := range extensions {
		if strings.HasSuffix(fileName, extension) {
			return true
		}
	}
	return false
}

func (s *FilesService) MIMETypeAllowed(mimeType string, allowedTypes []string) bool {
	return slices.Contains(allowedTypes, mimeType)
}
