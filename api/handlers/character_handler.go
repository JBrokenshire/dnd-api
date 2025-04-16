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

type CharacterHandler struct {
	server *api.Server
}

func NewCharacterHandler(server *api.Server) *CharacterHandler {
	return &CharacterHandler{
		server: server,
	}
}

// List godoc
// @Summary List characters
// @Description List characters (paginated)
// @ID characters-list
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param search query string false "Search characters by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.CharacterPaginatedResponse
// @Router /characters [get]
func (h *CharacterHandler) List(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)

	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	characters, page, pageSize := h.server.Repos.Character.GetCharacters(c, currentUser.ID, scopes)
	count := h.server.Repos.Character.CountCharacters(currentUser.ID, scopes)

	res := responses.NewCharacterPaginatedResponse(characters, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get character by ID
// @Description Get character by ID
// @ID characters-get
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} responses.CharacterResponse
// @Router /characters/{id} [get]
func (h *CharacterHandler) Get(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)
	id := c.Param("id")

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}

// Create godoc
// @Summary Create character
// @Description Create character
// @ID characters-create
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateCharacterRequest true "Character information"
// @Success 201 {object} responses.CharacterResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters [post]
func (h *CharacterHandler) Create(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)

	request := new(requests.CreateCharacterRequest)
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

	race := h.server.Repos.Race.GetById(request.RaceId)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	character := &models.Character{
		UserId:  currentUser.ID,
		Name:    request.Name,
		ClassId: class.ID,
		RaceId:  race.ID,
	}

	err := h.server.Repos.Character.Create(character)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong creating the character")
	}

	character.Class = *class
	character.Race = *race

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusCreated, res)
}

// Update godoc
// @Summary Update character
// @Description Update character
// @ID characters-update
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Param params body requests.UpdateCharacterRequest true "Character information"
// @Success 200 {object} responses.CharacterResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id} [put]
func (h *CharacterHandler) Update(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)
	id := c.Param("id")

	request := new(requests.UpdateCharacterRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	class := h.server.Repos.Class.GetById(request.ClassId)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	race := h.server.Repos.Race.GetById(request.RaceId)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	character.Name = request.Name
	character.ClassId = request.ClassId
	character.RaceId = request.RaceId
	character.Class = *class
	character.Race = *race

	err := h.server.Repos.Character.Update(character)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong updating the character")
	}

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}

// Delete godoc
// @Summary Delete character
// @Description Delete character
// @ID characters-delete
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id} [delete]
func (h *CharacterHandler) Delete(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)
	id := c.Param("id")

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	err := h.server.Repos.Character.Delete(character)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the character")
	}

	return responses.MessageResponse(c, http.StatusOK, "Character deleted successfully")
}
