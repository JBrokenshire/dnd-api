package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WeaponController struct {
	server.Server
}

func (w *WeaponController) GetAll(ctx echo.Context) error {
	weapons, err := w.Server.Stores.Weapon.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, weapons)
}
