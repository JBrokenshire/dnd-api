package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterSpellsController struct {
	server.Server
}

func (c *CharacterSpellsController) CharacterHasSpells(ctx echo.Context) error {
	characterID := ctx.Param("id")

	character, err := c.Server.Stores.Character.Get(characterID)
	if err != nil || character.ID == 0 {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	hasSpells, err := c.Server.Stores.CharacterSpells.GetHasSpellsByCharacterID(characterID)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, hasSpells)
}

func (c *CharacterSpellsController) GetCharacterSpells(ctx echo.Context) error {
	characterID := ctx.Param("id")

	character, err := c.Server.Stores.Character.Get(characterID)
	if err != nil || character.ID == 0 {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	spells, err := c.Server.Stores.CharacterSpells.GetSpellsByCharacterID(characterID)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, spells)
}

func (c *CharacterSpellsController) GetCharacterAttackSpells(ctx echo.Context) error {
	characterID := ctx.Param("id")

	character, err := c.Server.Stores.Character.Get(characterID)
	if err != nil || character.ID == 0 {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	attackSpells, err := c.Server.Stores.CharacterSpells.GetAttackSpellsByCharacterID(characterID)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, attackSpells)
}
