package controllers

import (
	"dnd-api/server"
	res "dnd-api/server/responses"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterInventoryController struct {
	server.Server
}

func (c *CharacterInventoryController) GetCharacterInventory(ctx echo.Context) error {
	characterInventory, err := c.Server.Stores.CharacterInventory.GetInventoryByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterInventory)
}

func (c *CharacterInventoryController) GetCharacterEquippedWeapons(ctx echo.Context) error {
	characterID := ctx.Param("id")

	character, err := c.Server.Stores.Character.Get(characterID)
	if err != nil || character.ID == 0 {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	characterEquippedWeapons, err := c.Server.Stores.CharacterInventory.GetEquippedWeaponsByCharacterID(characterID)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterEquippedWeapons)
}

func (c *CharacterInventoryController) GetCharacterEquippedArmour(ctx echo.Context) error {
	characterEquippedArmour, err := c.Server.Stores.CharacterInventory.GetEquippedArmourByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterEquippedArmour)
}

func (c *CharacterInventoryController) ToggleItemEquipped(ctx echo.Context) error {
	characterID := ctx.Param("characterID")
	itemID := ctx.Param("itemID")

	inventoryItem, err := c.Server.Stores.CharacterInventory.GetCharacterInventoryItemByID(characterID, itemID)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	if !inventoryItem.Item.Equippable {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("item can not be equipped"))
	}

	if inventoryItem.Location != "Equipment" {
		return ctx.JSON(http.StatusOK, inventoryItem)
	}

	inventoryItem.Equipped = !inventoryItem.Equipped
	err = c.Server.Stores.CharacterInventory.UpdateCharacterInventoryItem(inventoryItem)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, inventoryItem)
}
