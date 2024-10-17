package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterSkillAdvantage struct {
	ID           int    `gorm:"primary_key" json:"id"`
	CharacterID  int    `gorm:"not null" json:"character_id"`
	SkillName    string `gorm:"not null" json:"skill_name"`
	Advantage    bool   `gorm:"not null" json:"advantage"`
	Disadvantage bool   `gorm:"not null" json:"disadvantage"`
}

func (c *CharacterSkillAdvantage) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(skills, c.SkillName) {
		return fmt.Errorf("skill with name '%v' not found", c.SkillName)
	}

	return nil
}
