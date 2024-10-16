package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterInventoryItem struct {
	ID          int    `gorm:"primary_key" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	ItemID      int    `gorm:"not null" json:"item_id"`
	Equipped    bool   `json:"equipped"`
	Quantity    int    `json:"quantity"`
	Location    string `gorm:"not null" json:"location"`

	Item Item `json:"item"`
}

func (c *CharacterInventoryItem) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	var item Item
	err = db.Where("id = ?", c.ItemID).Find(&item).Error
	if err != nil {
		return fmt.Errorf("item with id '%v' not found", c.ItemID)
	}

	if !utils.SliceContains(locations, c.Location) {
		return fmt.Errorf("location '%s' is not valid", c.Location)
	}

	return nil
}
