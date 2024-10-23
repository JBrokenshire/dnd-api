package controllers

import (
	"dnd-api/db/models"
	"dnd-api/server"
	"dnd-api/server/requests"
	res "dnd-api/server/responses"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ClassController struct {
	server.Server
}

func (c *ClassController) GetAll(ctx echo.Context) error {
	classes, err := c.Server.Stores.Class.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, classes)
}

func (c *ClassController) Get(ctx echo.Context) error {
	class, err := c.Server.Stores.Class.Get(ctx.Param("id"))

	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, *class)
}

func (c *ClassController) Update(ctx echo.Context) error {
	existingClass, err := c.Server.Stores.Class.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	updatedClassRequest := new(requests.ClassRequest)
	if err := ctx.Bind(&updatedClassRequest); err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}
	if updatedClassRequest == nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid class request body"))
	}

	if updatedClassRequest.Name == "" {
		updatedClassRequest.Name = existingClass.Name
	}
	if updatedClassRequest.ShortDescription == "" {
		updatedClassRequest.ShortDescription = existingClass.ShortDescription
	}
	if updatedClassRequest.LongDescription == "" {
		updatedClassRequest.LongDescription = existingClass.LongDescription
	}

	updatedClass := &models.Class{
		ID:               existingClass.ID,
		Name:             updatedClassRequest.Name,
		ShortDescription: updatedClassRequest.ShortDescription,
		LongDescription:  updatedClassRequest.LongDescription,
	}

	err = c.Server.Stores.Class.Update(updatedClass)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, updatedClass)
}
