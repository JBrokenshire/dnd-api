package handlers

import (
	"dnd-api/api"
	"dnd-api/api/responses"
	"dnd-api/db/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type CharacterInspirationHandler struct {
	server *api.Server
}

func NewCharacterInspirationHandler(server *api.Server) *CharacterInspirationHandler {
	return &CharacterInspirationHandler{
		server: server,
	}
}

// Update godoc
// @Summary Update character inspiration
// @Description Update character inspiration
// @ID characters-update-inspiration
// @Tags Character Inspiration Actions
// @Param id path string true "Character ID"
// @Success 200 {object} responses.CharacterResponse
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id}/inspiration [put]
func (h *CharacterInspirationHandler) Update(c echo.Context) error {
	id := c.Param("id")
	currentUser := c.Get("currentUser").(*models.User)

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	character.Inspiration = !character.Inspiration
	err := h.server.Repos.Character.Update(character)
	if err != nil {
		log.Printf("Error toggling character inspiration: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error toggling character inspiration")
	}

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}
