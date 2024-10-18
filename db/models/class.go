package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Class struct {
	ID                  int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                string  `json:"name"`
	ShortDescription    string  `json:"short_description"`
	LongDescription     string  `json:"long_description"`
	SpellcastingAbility *string `json:"spellcasting_ability"`
}

func (c *Class) BeforeCreate(_ *gorm.DB) error {
	if c.SpellcastingAbility != nil {
		if !utils.SliceContains(spellcastingAbilities, *c.SpellcastingAbility) {
			return fmt.Errorf("spellcasting ability %q is not valid", c.SpellcastingAbility)
		}
	}

	return nil
}
