package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterConditionsController struct {
	server.Server
}

func (c *CharacterConditionsController) GetCharacterConditions(ctx echo.Context) error {
	characterConditions, err := c.Server.Stores.CharacterConditions.GetConditionsByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterConditions)
}
