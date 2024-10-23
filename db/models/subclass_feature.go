package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type SubclassFeature struct {
	ID         int `gorm:"primary_key" json:"id"`
	SubclassID int `json:"subclass_id"`
	FeatureID  int `json:"feature_id"`
	Level      int `json:"level"`

	Feature Feature `json:"feature"`
}

func (s *SubclassFeature) BeforeCreate(db *gorm.DB) error {
	var subclass Subclass
	err := db.Where("id = ?", s.SubclassID).First(&subclass).Error
	if err != nil {
		return fmt.Errorf("error getting subclass with id %v - %v", s.SubclassID, err)
	}

	var feature Feature
	err = db.Where("id = ?", s.FeatureID).First(&feature).Error
	if err != nil {
		return fmt.Errorf("error getting feature with id %v - %v", s.FeatureID, err)
	}

	if s.Level == 0 {
		s.Level = 1
	}

	return nil
}
