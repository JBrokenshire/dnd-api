package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Subclass struct {
	ID      int    `gorm:"primary_key" json:"id"`
	ClassID int    `json:"class_id"`
	Name    string `json:"name"`
}

func (s *Subclass) BeforeCreate(db *gorm.DB) error {
	var class Class
	err := db.Where("id = ?", s.ClassID).First(&class, s.ClassID).Error
	if err != nil {
		return fmt.Errorf("class with id %v not found - %v", s.ClassID, err)
	}

	return nil
}
