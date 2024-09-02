package controller

import (
	"dungeons-and-dragons/db/models"
	"dungeons-and-dragons/db/stores"
	res "dungeons-and-dragons/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ClassController struct {
	Store stores.ClassStore
}

func (c *ClassController) GetAll(ctx echo.Context) error {
	classes, err := c.Store.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, classes)
}

func (c *ClassController) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}

	class, err := c.Store.Get(id)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, *class)
}

func (c *ClassController) Update(ctx echo.Context) error {
	updatedClass := new(models.Class)
	if err := ctx.Bind(&updatedClass); err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}
	updatedClass.ID = id

	err = c.Store.Update(updatedClass)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, updatedClass)
}