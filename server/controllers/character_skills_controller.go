package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterSkillsController struct {
	server.Server
}

func (c *CharacterSkillsController) GetProficientByCharacterID(ctx echo.Context) error {
	characterProficientSkills, err := c.Server.Stores.CharacterSkills.GetProficientByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterProficientSkills)
}

func (c *CharacterSkillsController) GetAdvantagesByCharacterID(ctx echo.Context) error {
	characterID := ctx.Param("id")

	character, err := c.Server.Stores.Character.Get(characterID)
	if err != nil || character.ID == 0 {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	characterSkillsAdvantages, err := c.Server.Stores.CharacterSkillsAdvantages.GetSkillsAdvantagesByCharacterID(characterID)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, characterSkillsAdvantages)
}
