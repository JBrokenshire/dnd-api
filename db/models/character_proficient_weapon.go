package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterProficientWeapon struct {
	ID          int    `gorm:"autoIncrement;primary_key" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	Weapon      string `gorm:"not null" json:"weapon"`
}

func (c *CharacterProficientWeapon) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(weaponTypes, c.Weapon) {
		return fmt.Errorf("weapon '%s' is not valid", c.Weapon)
	}

	return nil
}
