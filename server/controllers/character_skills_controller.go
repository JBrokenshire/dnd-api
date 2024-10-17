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
