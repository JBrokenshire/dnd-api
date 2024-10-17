package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Armour struct {
	ItemID               int    `gorm:"primary_key" json:"item_id"`
	Type                 string `json:"type"`
	BaseAC               int    `json:"base_ac"`
	MaxDexterityModifier int    `json:"max_dex_modifier"`
	StrengthRequirement  int    `json:"strength_requirement"`
	StealthDisadvantage  bool   `json:"stealth_disadvantage"`

	Item Item `json:"item"`
}

func (a *Armour) BeforeCreate(db *gorm.DB) error {
	var item Item
	err := db.Where("id = ?", a.ItemID).First(&item).Error
	if err != nil {
		return fmt.Errorf("item with id %v not found", a.ItemID)
	}

	if !utils.SliceContains(armourTypes, a.Type) {
		return fmt.Errorf("armour type %v is not valid", a.ItemID)
	}

	return nil
}
