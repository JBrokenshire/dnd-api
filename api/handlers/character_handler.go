package handlers

import (
	s "dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	m "dnd-api/db/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type CharacterHandler struct {
	server *s.Server
}

func NewCharacterHandler(server *s.Server) *CharacterHandler {
	return &CharacterHandler{server: server}
}

// List godoc
// @Summary List characters
// @Description Get a paginated list of characters
// @ID characters-list
// @Tags Character
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Number of items to return"
// @Success 200 {object} responses.CharacterPaginatedResponse
// @Router /characters [get]
func (h *CharacterHandler) List(c echo.Context) error {
	characters, page, pageSize := h.server.Repos.Character.GetCharacters(c)
	count := h.server.Repos.Character.Count()

	res := responses.NewCharacterPaginatedResponse(characters, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get character
// @Description Get a character by their ID
// @ID characters-get
// @Tags Character
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} responses.CharacterResponse
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id} [get]
func (h *CharacterHandler) Get(c echo.Context) error {
	id := c.Param("id")

	character := h.server.Repos.Character.GetCharacter(id)
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
// @Tags Character
// @Accept json
// @Produce json
// @Param params body requests.CreateCharacterRequest true "Character information"
// @Success 200 {object} responses.CharacterResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters [post]
func (h *CharacterHandler) Create(c echo.Context) error {
	request := new(requests.CreateCharacterRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty or not valid: %v", err.Error()))
	}

	class := h.server.Repos.Class.GetClass(request.ClassId)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	race := h.server.Repos.Race.GetRace(request.RaceId)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	character := &m.Character{
		Name:    request.Name,
		ClassId: request.ClassId,
		RaceId:  request.RaceId,
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
// @Tags Character
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
	id := c.Param("id")

	request := new(requests.UpdateCharacterRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty or not valid: %v", err.Error()))
	}

	character := h.server.Repos.Character.GetCharacter(id)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	class := h.server.Repos.Class.GetClass(request.ClassId)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	race := h.server.Repos.Race.GetRace(request.RaceId)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	character.Name = request.Name
	character.ClassId = request.ClassId
	character.RaceId = request.RaceId

	err := h.server.Repos.Character.Update(&character)
	if err != nil {
		log.Printf("Error updating character: %v", err)
		return responses.Response(c, http.StatusInternalServerError, "Something went wrong updating the character")
	}

	character.Race = *race
	character.Class = *class

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}

// Delete godoc
// @Summary Delete character
// @Description Delete character
// @ID characters-delete
// @Tags Class
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id} [delete]
func (h *CharacterHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	character := h.server.Repos.Character.GetCharacter(id)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	err := h.server.Repos.Character.Delete(character)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the character")
	}

	return responses.MessageResponse(c, http.StatusOK, "Character deleted successfully")
}
