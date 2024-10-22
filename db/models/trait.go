package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Trait struct {
	ID          int     `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Action      *string `json:"action"`
	ActionType  *string `json:"action_type"`
	ActionUses  *int    `json:"action_uses"`
	ActionReset *string `json:"action_reset"`
}

func (t *Trait) BeforeCreate(_ *gorm.DB) error {
	if t.ActionReset != nil {
		if !utils.SliceContains(rests, *t.ActionReset) {
			return fmt.Errorf("action reset %q is not valid", *t.ActionReset)
		}
	}

	if t.ActionType != nil {
		if !utils.SliceContains(actionTypes, *t.ActionType) {
			return fmt.Errorf("action type %q is not valid", *t.ActionType)
		}
	}

	return nil
}
