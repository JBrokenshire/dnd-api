package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
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
		Name: request.Name,
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

	request := new(requests.CreateClassRequest)
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
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /classes/{id} [delete]
func (h *ClassHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	class := h.server.Repos.Class.GetById(id)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	err := h.server.Repos.Class.Delete(class)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the class")
	}

	return responses.MessageResponse(c, http.StatusOK, "Class deleted successfully")
}
