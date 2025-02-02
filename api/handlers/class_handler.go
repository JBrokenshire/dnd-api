package handlers

import (
	s "github.com/JBrokenshire/dnd-api/api"
	"github.com/JBrokenshire/dnd-api/api/requests"
	"github.com/JBrokenshire/dnd-api/api/responses"
	m "github.com/JBrokenshire/dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ClassHandler struct {
	server *s.Server
}

func NewClassHandler(server *s.Server) *ClassHandler {
	return &ClassHandler{server: server}
}

// List godoc
// @Summary List classes
// @Description List classes (paginated)
// @ID class-list
// @Tags Class
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.ClassPaginatedResponse
// @Router /classes [get]
func (h *ClassHandler) List(c echo.Context) error {
	var scopes []func(db *gorm.DB) *gorm.DB

	classes, page, pageSize := h.server.Repos.Class.GetClasses(c, scopes)
	count := h.server.Repos.Class.Count(scopes)

	res := responses.NewClassPaginatedResponse(classes, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get class
// @Description Get class by ID
// @ID class-get
// @Tags Class
// @Param id query int true "Class ID"
// @Success 200 {object} responses.ClassResponse
// @Failure 404 {object} responses.Error
// @Router /classes [get]
func (h *ClassHandler) Get(c echo.Context) error {
	id := c.Param("id")

	var class m.Class
	h.server.Repos.Class.GetById(id, &class)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Cannot find class")
	}

	res := responses.NewClassResponse(&class)
	return responses.Response(c, http.StatusOK, res)
}

// Create godoc
// @Summary Create class
// @Description Create class
// @ID class-create
// @Tags Class
// @Accept json
// @Produce json
// @Param params body requests.ClassCreateRequest true "Class information"
// @Success 201 {object} responses.ClassResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /classes [post]
func (h *ClassHandler) Create(c echo.Context) error {
	request := requests.ClassCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	class := &m.Class{
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
		LongDescription:  request.LongDescription,
	}
	err := h.server.Repos.Class.Create(class)
	if err != nil {
		return responses.Response(c, http.StatusInternalServerError, "Error creating class: "+err.Error())
	}

	res := responses.NewClassResponse(class)
	return responses.Response(c, http.StatusCreated, res)
}
