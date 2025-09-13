package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/db/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type CharacterUsedSpellSlotHandler struct {
	server *api.Server
}

func NewCharacterUsedSpellSlotHandler(server *api.Server) *CharacterUsedSpellSlotHandler {
	return &CharacterUsedSpellSlotHandler{
		server: server,
	}
}

// Update godoc
// @Summary Update character used spell slots
// @Description Update character used spell slots
// @ID characters-used-spell-slots-update
// @Tags Character Spells Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Param spellLevel path int true "Spell Level"
// @Param params body requests.UpdateCharacterUsedSpellSlotRequest true "Character used spell slot information"
// @Success 200 {object} responses.CharacterResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id}/spell-slots/{spellLevel} [put]
func (h *CharacterUsedSpellSlotHandler) Update(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)
	id := c.Param("id")
	spellLevel := c.Param("spellLevel")

	request := new(requests.UpdateCharacterUsedSpellSlotRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}
	spellLevelInt, err := strconv.Atoi(spellLevel)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Spell level must be an integer")
	}
	if spellLevelInt < 1 || spellLevelInt > 9 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Spell level must be between 1 and 9")
	}

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	usedSpellSlot := h.server.Repos.CharacterUsedSpellSlot.Get(character.ID, spellLevel)
	if usedSpellSlot.ID == 0 {
		usedSpellSlot = &models.CharacterUsedSpellSlot{
			CharacterId: character.ID,
			SpellLevel:  spellLevelInt,
		}

		err := h.server.Repos.CharacterUsedSpellSlot.Create(usedSpellSlot)
		if err != nil {
			log.Println("Error creating new entry for character used spell slot: ", err.Error())
			return responses.ErrorResponse(c, http.StatusInternalServerError, "Error creating new entry for character used spell slot")
		}
	}

	usedSpellSlot.SpellSlotsUsed = request.SpellSlotsUsed
	err = h.server.Repos.CharacterUsedSpellSlot.Update(usedSpellSlot)
	if err != nil {
		log.Println("Error updating character used spell slot: ", err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error updating character used spell slot.")
	}

	return responses.MessageResponse(c, http.StatusOK, "Character used spell slots updated")
}
