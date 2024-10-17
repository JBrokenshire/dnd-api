package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ItemController struct {
	server.Server
}

func (c *ItemController) GetAll(ctx echo.Context) error {
	items, err := c.Server.Stores.Item.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, items)
}
