package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/db/models"
	"dnd-api/pkg/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type CharacterInventoryItemHandler struct {
	server *api.Server
}

func NewCharacterInventoryItemHandler(server *api.Server) *CharacterInventoryItemHandler {
	return &CharacterInventoryItemHandler{
		server: server,
	}
}

// Update godoc
// @Summary Update character inventory item
// @Description Update character inventory item
// @ID characters-update-inventory-item
// @Tags Character Inventory Item Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Param inventoryItemId path int true "Inventory Item ID"
// @Param params body requests.UpdateCharacterInventoryItemRequest true "Inventory Item information"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id}/inventory-item/{inventoryItemId} [put]
func (h *CharacterInventoryItemHandler) Update(c echo.Context) error {
	characterId := c.Param("id")
	inventoryItemId := c.Param("inventoryItemId")
	currentUser := c.Get("currentUser").(*models.User)

	request := new(requests.UpdateCharacterInventoryItemRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	character := h.server.Repos.Character.GetById(characterId, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	inventoryItem := h.server.Repos.CharacterInventoryItem.GetById(characterId, inventoryItemId)
	if inventoryItem.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Inventory item not found")
	}

	inventoryItem.Equipped = utils.BoolPointer(request.Equipped)
	err := h.server.Repos.CharacterInventoryItem.Update(inventoryItem)
	if err != nil {
		log.Printf("Error updating character health: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error updating character health")
	}

	return responses.MessageResponse(c, http.StatusOK, "Character inventory item updated")
}
