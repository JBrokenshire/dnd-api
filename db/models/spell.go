package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Spell struct {
	ID          int     `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	Level       int     `json:"level"`
	CastingTime string  `json:"casting_time"`
	Distance    string  `json:"range"`
	Effect      *string `json:"effect"`
	Damage      *string `json:"damage"`
	DamageType  *string `json:"damage_type"`
	Save        *string `json:"save"`
	Notes       string  `json:"notes"`
}

func (s *Spell) BeforeCreate(_ *gorm.DB) error {
	if s.Save != nil {
		if !utils.SliceContains(abilityAbbreviations, *s.Save) {
			return fmt.Errorf("spell save %q is not valid", *s.Save)
		}
	}

	if s.DamageType != nil {
		if !utils.SliceContains(damageTypes, *s.DamageType) {
			return fmt.Errorf("spell damage type %q is not valid", *s.DamageType)
		}
	}

	return nil
}
