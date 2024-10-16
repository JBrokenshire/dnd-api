package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Item struct {
	ID         int     `gorm:"primary_key" json:"id"`
	Name       string  `gorm:"not null" json:"name"`
	Meta       string  `gorm:"not null" json:"meta"`
	Weight     float64 `json:"weight"`
	Cost       float64 `json:"cost"`
	Notes      string  `json:"notes"`
	Rarity     string  `gorm:"not null" json:"rarity"`
	Equippable bool    `gorm:"not null" json:"equippable"`
}

func (i *Item) BeforeCreate(_ *gorm.DB) error {
	if i.Rarity == "" {
		i.Rarity = "Common"
	}

	if !utils.SliceContains(rarities, i.Rarity) {
		return fmt.Errorf("rarity '%s' is not valid", i.Rarity)
	}

	return nil
}
