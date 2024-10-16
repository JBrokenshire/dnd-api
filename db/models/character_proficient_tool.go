package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterProficientTool struct {
	ID          int    `gorm:"primary_key auto_increment" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	Tool        string `gorm:"not null" json:"tool"`
}

func (c *CharacterProficientTool) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(tools, c.Tool) {
		return fmt.Errorf("tool '%s' is not valid", c.Tool)
	}

	return nil
}
