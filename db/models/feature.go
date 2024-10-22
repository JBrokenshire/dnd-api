package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Feature struct {
	ID          int     `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Action      *string `json:"action"`
	ActionType  *string `json:"action_type"`
	ActionUses  *int    `json:"action_uses"`
	ActionReset *string `json:"action_reset"`
}

func (f *Feature) BeforeCreate(_ *gorm.DB) error {
	if f.ActionReset != nil {
		if !utils.SliceContains(rests, *f.ActionReset) {
			return fmt.Errorf("action reset %q is not valid", *f.ActionReset)
		}
	}

	if f.ActionType != nil {
		if !utils.SliceContains(actionTypes, *f.ActionType) {
			return fmt.Errorf("action type %q is not valid", *f.ActionType)
		}
	}

	return nil
}
