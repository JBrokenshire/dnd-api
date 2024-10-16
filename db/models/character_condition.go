package models

import (
	"dnd-api/utils"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterCondition struct {
	ID            int    `gorm:"primary_key;auto_increment" json:"id"`
	CharacterID   int    `gorm:"not null" json:"character_id"`
	ConditionName string `gorm:"not null" json:"condition_name"`
}

func (c *CharacterCondition) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(conditions, c.ConditionName) {
		return errors.New(fmt.Sprintf("condition '%s' is not valid", c.ConditionName))
	}

	return nil
}
