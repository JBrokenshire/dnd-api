package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Shield struct {
	ItemID  int `gorm:"primary_key" json:"item_id"`
	BonusAC int `json:"bonus_ac"`

	Item Item `json:"item"`
}

func (s *Shield) BeforeCreate(db *gorm.DB) error {
	var item Item
	err := db.Where("id = ?", s.ItemID).First(&item).Error
	if err != nil {
		return fmt.Errorf("item with id %v not found", s.ItemID)
	}

	return nil
}
