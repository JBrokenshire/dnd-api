package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterDefense struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	DamageType  string `gorm:"not null" json:"damage_type"`
	DefenseType string `gorm:"not null" json:"defense_type"`
}

func (c *CharacterDefense) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(defenseTypes, c.DefenseType) {
		return fmt.Errorf("defense type '%s' is not valid", c.DefenseType)
	}

	return nil
}
