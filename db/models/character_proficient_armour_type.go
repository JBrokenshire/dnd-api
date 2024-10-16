package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterProficientArmourType struct {
	ID          int    `gorm:"autoIncrement;primary_key" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	ArmourType  string `gorm:"not null" json:"armour_type"`
}

func (c *CharacterProficientArmourType) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(armourTypes, c.ArmourType) {
		return fmt.Errorf("armour type '%s' is not valid", c.ArmourType)
	}

	return nil
}
