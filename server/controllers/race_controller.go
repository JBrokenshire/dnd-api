package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RaceController struct {
	server.Server
}

func (r *RaceController) GetAll(ctx echo.Context) error {
	races, err := r.Server.Stores.Race.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, races)
}

func (r *RaceController) Get(ctx echo.Context) error {
	race, err := r.Server.Stores.Race.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, race)
}

func (r *RaceController) GetTraits(ctx echo.Context) error {
	raceID := ctx.Param("id")

	race, err := r.Server.Stores.Race.Get(raceID)
	if err != nil || race.ID == 0 {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	traits, err := r.Server.Stores.Race.GetTraits(raceID)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, traits)
}
