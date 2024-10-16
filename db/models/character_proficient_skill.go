package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterProficientSkill struct {
	ID              int    `gorm:"autoIncrement;primary_key" json:"id"`
	CharacterID     int    `gorm:"not null" json:"character_id"`
	SkillName       string `gorm:"not null" json:"skill_name"`
	ProficiencyType string `gorm:"not null default:'Proficiency'" json:"proficiency_type"`
}

func (c *CharacterProficientSkill) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(proficiencyTypes, c.ProficiencyType) {
		return fmt.Errorf("proficiency type '%s' is not valid", c.ProficiencyType)
	}

	if !utils.SliceContains(skills, c.SkillName) {
		return fmt.Errorf("skill name '%s' is not valid", c.SkillName)
	}

	return nil
}
