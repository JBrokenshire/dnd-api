package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterMoneyController struct {
	server.Server
}

func (c *CharacterMoneyController) GetCharacterMoney(ctx echo.Context) error {
	characterMoney, err := c.Server.Stores.CharacterMoney.GetMoneyByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterMoney)
}
