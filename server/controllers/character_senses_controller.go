package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterSensesController struct {
	server.Server
}

func (c *CharacterSensesController) GetSensesByCharacterID(ctx echo.Context) error {
	characterSenses, err := c.Server.Stores.CharacterSenses.GetSensesByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterSenses)
}
