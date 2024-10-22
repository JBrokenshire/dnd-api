package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type RaceTrait struct {
	ID      int `gorm:"primary_key" json:"id"`
	RaceID  int `json:"race_id"`
	TraitID int `json:"trait_id"`

	Trait Trait `json:"trait"`
}

func (r *RaceTrait) BeforeCreate(db *gorm.DB) error {
	var race Race
	err := db.Where("id = ?", r.RaceID).First(&race).Error
	if err != nil {
		return fmt.Errorf("error getting race with id: %v - %v", r.RaceID, err)
	}

	var trait Trait
	err = db.Where("id = ?", r.TraitID).First(&trait).Error
	if err != nil {
		return fmt.Errorf("error getting trait with id: %v - %v", r.TraitID, err)
	}

	return nil
}
