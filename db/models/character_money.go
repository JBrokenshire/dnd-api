package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterMoney struct {
	ID          int    `gorm:"primary_key" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	Money       string `gorm:"not null" json:"money"`
	Amount      int    `gorm:"not null" json:"amount"`
}

func (c *CharacterMoney) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(moneyTypes, c.Money) {
		return fmt.Errorf("money type '%s' is not valid", c.Money)
	}

	if c.Amount < 0 {
		return fmt.Errorf("money value '%d' can't be negative", c.Amount)
	}

	return nil
}
