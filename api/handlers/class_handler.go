package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/db/models"
	"dnd-api/pkg/closer"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"log"
	"net/http"
	"path/filepath"
)

type ClassHandler struct {
	server *api.Server
}

func NewClassHandler(server *api.Server) *ClassHandler {
	return &ClassHandler{
		server: server,
	}
}

// List godoc
// @Summary List classes
// @Description List classes (paginated)
// @ID classes-list
// @Tags Class Actions
// @Accept json
// @Produce json
// @Param search query string false "Search classes by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.ClassPaginatedResponse
// @Router /classes [get]
func (h *ClassHandler) List(c echo.Context) error {
	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	classes, page, pageSize := h.server.Repos.Class.GetClasses(c, scopes)
	count := h.server.Repos.Class.CountClasses(scopes)

	res := responses.NewClassPaginatedResponse(classes, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get class by ID
// @Description Get class by ID
// @ID classes-get
// @Tags Class Actions
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} responses.ClassResponse
// @Failure 404 {object} responses.Error
// @Router /classes/{id} [get]
func (h *ClassHandler) Get(c echo.Context) error {
	id := c.Param("id")

	class := h.server.Repos.Class.GetById(id)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	res := responses.NewClassResponse(class)
	return responses.Response(c, http.StatusOK, res)
}

// Create godoc
// @Summary Create class
// @Description Create class
// @ID classes-create
// @Tags Class Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateClassRequest true "Class information"
// @Success 201 {object} responses.ClassResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /classes [post]
func (h *ClassHandler) Create(c echo.Context) error {
	request := new(requests.CreateClassRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	class := &models.Class{
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
		PrimaryAbility:   request.PrimaryAbility,
		HitPointDieValue: request.HitPointDieValue,
		Saves:            request.Saves,
	}
	err := h.server.Repos.Class.Create(class)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong creating the class")
	}

	res := responses.NewClassResponse(class)
	return responses.Response(c, http.StatusCreated, res)
}

// Update godoc
// @Summary Update class
// @Description Update class
// @ID classes-update
// @Tags Class Actions
// @Accept json
// @Produce json
// @Param id path string true "Class ID"
// @Param params body requests.UpdateClassRequest true "Class information"
// @Success 200 {object} responses.ClassResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /classes/{id} [put]
func (h *ClassHandler) Update(c echo.Context) error {
	id := c.Param("id")

	request := new(requests.UpdateClassRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	class := h.server.Repos.Class.GetById(id)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	class.Name = request.Name
	class.ShortDescription = request.ShortDescription
	class.PrimaryAbility = request.PrimaryAbility
	class.HitPointDieValue = request.HitPointDieValue
	class.Saves = request.Saves

	err := h.server.Repos.Class.Update(class)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong updating the class")
	}

	res := responses.NewClassResponse(class)
	return responses.Response(c, http.StatusOK, res)
}

// Delete godoc
// @Summary Delete class
// @Description Delete class
// @ID classes-delete
// @Tags Class Actions
// @Accept json
// @Produce json
// @Param id path string true "Class ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /classes/{id} [delete]
func (h *ClassHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	class := h.server.Repos.Class.GetById(id)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	subclasses := h.server.Repos.Class.GetSubclasses(class.ID)
	if len(subclasses) > 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Can't delete class with subclasses")
	}

	// Delete logo
	if class.Logo.ID != 0 {
		// Delete from file store
		err := h.server.Dependencies.GetFileStore().Delete(fmt.Sprintf("%v/%v", class.Logo.FileLocation, class.Logo.Filename))
		if err != nil {
			log.Println("Error deleting class logo from file store: ", err.Error())
			return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the class logo from the file store")
		}

		err = h.server.Repos.File.Delete(class.Logo)
		if err != nil {
			log.Println("Error deleting class logo from database: ", err.Error())
			return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the class logo from the database")
		}
	}

	err := h.server.Repos.Class.Delete(class)
	if err != nil {
		log.Println("Error deleting class from the database: ", err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the class")
	}

	return responses.MessageResponse(c, http.StatusOK, "Class deleted successfully")
}

// UploadLogo godoc
// @Summary Upload class logo
// @Description Upload class logo
// @ID classes-upload-logo
// @Tags Class File Actions
// @Accept json
// @Produce json
// @Param id path string true "Class ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /classes/{id}/upload/logo [post]
func (h *ClassHandler) UploadLogo(c echo.Context) error {
	id := c.Param("id")

	class := h.server.Repos.Class.GetById(id)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
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

	path := fmt.Sprintf("classes/%v", class.ID)
	err = h.server.Dependencies.GetFileStore().Save(path, src, file.Filename)
	if err != nil {
		log.Printf("Error saving file: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error saving file")
	}

	// Create DB record
	logoFile := &models.File{
		Model:        models.FileModelClassLogo,
		ModelId:      class.ID,
		Filename:     newFilename,
		FileLocation: path,
	}
	if err := h.server.Repos.File.Create(logoFile); err != nil {
		log.Printf("Error creating file record: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error creating file record")
	}

	return responses.MessageResponse(c, http.StatusOK, "File uploaded")
}
