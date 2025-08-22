package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/db/models"
	"github.com/labstack/echo/v4"
	"log"
	"math"
	"net/http"
)

type CharacterHealthHandler struct {
	server *api.Server
}

func NewCharacterHealthHandler(server *api.Server) *CharacterHealthHandler {
	return &CharacterHealthHandler{
		server: server,
	}
}

// Update godoc
// @Summary Update character health
// @Description Update character health
// @ID characters-update-health
// @Tags Character Health Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Param params body requests.UpdateCharacterHealthRequest true "Health information"
// @Success 200 {object} responses.CharacterResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id}/health [put]
func (h *CharacterHealthHandler) Update(c echo.Context) error {
	id := c.Param("id")
	currentUser := c.Get("currentUser").(*models.User)

	request := new(requests.UpdateCharacterHealthRequest)
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

	character.CurrentHitPoints = uint(math.Max(math.Min(float64(request.CurrentHitPoints), float64(character.MaxHitPoints)), 0))
	err := h.server.Repos.Character.Update(character)
	if err != nil {
		log.Printf("Error updating character health: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error updating character health")
	}

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}
