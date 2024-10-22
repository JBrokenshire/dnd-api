package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SubclassController struct {
	server.Server
}

func (c *SubclassController) GetFeatures(ctx echo.Context) error {
	subclassID := ctx.Param("id")

	subclass, err := c.Server.Stores.Subclass.Get(subclassID)
	if err != nil || subclass.ID == 0 {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	features, err := c.Server.Stores.Subclass.GetFeatures(subclass.ID)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, features)
}
