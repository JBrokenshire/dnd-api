package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	m "dnd-api/db/models"
	"dnd-api/pkg/closer"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"log"
	"net/http"
	"path/filepath"
)

type SubclassHandler struct {
	server *api.Server
}

func NewSubclassHandler(server *api.Server) *SubclassHandler {
	return &SubclassHandler{server: server}
}

// List godoc
// @Summary List subclasses
// @Description List subclasses for class (paginated)
// @ID subclasses-list
// @Tags Subclass Actions
// @Accept json
// @Produce json
// @Param classId path string true "Class ID"
// @Param search query string false "Search subclasses by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.SubclassPaginatedResponse
// @Failure 404 {object} responses.Error
// @Router /subclasses/{classId} [get]
func (h *SubclassHandler) List(c echo.Context) error {
	id := c.Param("classId")

	class := h.server.Repos.Class.GetById(id)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	subclasses, page, pageSize := h.server.Repos.Subclass.GetSubclasses(c, class.ID, scopes)
	count := h.server.Repos.Subclass.CountSubclasses(class.ID, scopes)

	res := responses.NewSubclassPaginatedResponse(subclasses, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get subclass by ID
// @Description Get subclass by ID
// @ID subclasses-get
// @Tags Class Actions
// @Accept json
// @Produce json
// @Param classId path string true "Class ID"
// @Param subclassId path string true "Subclass ID"
// @Success 200 {object} responses.SubclassResponse
// @Failure 404 {object} responses.Error
// @Router /subclasses/{classId}/{subclassId} [get]
func (h *SubclassHandler) Get(c echo.Context) error {
	classId := c.Param("classId")
	subclassId := c.Param("subclassId")

	subclass := h.server.Repos.Subclass.GetById(subclassId, classId)
	if subclass.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Subclass not found")
	}

	res := responses.NewSubclassResponse(subclass)
	return responses.Response(c, http.StatusOK, res)
}

// Create godoc
// @Summary Create subclass
// @Description Create subclass
// @ID subclasses-create
// @Tags Subclass Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateSubclassRequest true "Subclass information"
// @Success 201 {object} responses.SubclassResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /subclasses [post]
func (h *SubclassHandler) Create(c echo.Context) error {
	request := new(requests.CreateSubclassRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	class := h.server.Repos.Class.GetById(request.ClassId)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	subclass := &m.Subclass{
		ClassId:          request.ClassId,
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
	}

	err := h.server.Repos.Subclass.Create(subclass)
	if err != nil {
		log.Println("Error creating subclass: ", err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong creating the subclass")
	}

	res := responses.NewSubclassResponse(subclass)
	return responses.Response(c, http.StatusCreated, res)
}

// UploadLogo godoc
// @Summary Upload subclass logo
// @Description Upload subclass logo
// @ID subclasses-upload-logo
// @Tags Subclass File Actions
// @Accept json
// @Produce json
// @Param classId path string true "Class ID"
// @Param subclassId path int true "Subclass ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /subclasses/{classId}/{subclassId}/upload/logo [post]
func (h *SubclassHandler) UploadLogo(c echo.Context) error {
	classId := c.Param("classId")
	subclassId := c.Param("subclassId")

	class := h.server.Repos.Class.GetById(classId)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	subclass := h.server.Repos.Subclass.GetById(subclassId, classId)
	if subclass.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Subclass not found")
	}

	// Check the file mimetype - We only want to accept images
	fileName, fileMimeType, err := h.server.Dependencies.GetFileService().ReadFileInfo(c)
	if err != nil {
		log.Printf("Error reading file info: %v", err)
		return responses.ErrorResponse(c, http.StatusBadRequest, "Unable to read file")
	}
	fileExtension := filepath.Ext(fileName)
	allowedFileExtensions := []string{"jpg", "jpeg", "png", "webp"}
	if !h.server.Dependencies.GetFileService().FileExtensionAllowed(fileExtension, allowedFileExtensions) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Invalid file type")
	}

	allowedMIMEs := []string{"image/jpeg", "image/png", "image/webp"}
	if !h.server.Dependencies.GetFileService().MIMETypeAllowed(fileMimeType, allowedMIMEs) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Invalid file content type")
	}

	// Get file from request
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error getting file from form: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error getting file from request")
	}

	// Change the file name
	newFilename := random.String(32)
	if fileExtension != "" {
		newFilename += fileExtension
	}
	file.Filename = newFilename

	// Upload file
	src, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error opening file")
	}
	defer closer.Close(src)

	path := fmt.Sprintf("classes/%v/subclasses/%v", class.ID, subclass.ID)
	err = h.server.Dependencies.GetFileStore().Save(path, src, file.Filename)
	if err != nil {
		log.Printf("Error saving file: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error saving file")
	}

	// Create DB record
	logoFile := &m.File{
		Model:        m.FileModelSubclassLogo,
		ModelId:      subclass.ID,
		Filename:     newFilename,
		FileLocation: path,
	}
	if err := h.server.Repos.File.Create(logoFile); err != nil {
		log.Printf("Error creating file record: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error creating file record")
	}

	return responses.MessageResponse(c, http.StatusOK, "File uploaded")
}
