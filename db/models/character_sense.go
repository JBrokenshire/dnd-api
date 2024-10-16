package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterSense struct {
	ID          int    `gorm:"autoIncrement;primary_key" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	SenseName   string `gorm:"not null" json:"sense_name"`
	Distance    int    `gorm:"not null" json:"distance"`
}

func (c *CharacterSense) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(senses, c.SenseName) {
		return fmt.Errorf("proficiency type '%s' is not valid", c.SenseName)
	}

	if c.Distance <= 0 {
		return fmt.Errorf("distance needs to be a positive number")
	}

	return nil
}
