package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type ClassFeature struct {
	ID        int `gorm:"primary_key" json:"id"`
	ClassID   int `json:"class_id"`
	FeatureID int `json:"feature_id"`
	Level     int `json:"level"`

	Feature Feature `json:"feature"`
}

func (c *ClassFeature) BeforeCreate(db *gorm.DB) error {
	var class Class
	err := db.Where("id = ?", c.ClassID).First(&class).Error
	if err != nil {
		return fmt.Errorf("class with id %v could not be found", c.ClassID)
	}

	var feature Feature
	err = db.Where("id = ?", c.FeatureID).First(&feature).Error
	if err != nil {
		return fmt.Errorf("feature with id %v could not be found", c.FeatureID)
	}

	if c.Level == 0 {
		c.Level = 1
	}

	return nil
}
