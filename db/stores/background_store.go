package stores

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

type BackgroundStore interface {
	IsValidName(name string) bool
}

type GormBackgroundStore struct {
	DB *gorm.DB
}

func NewGormBackgroundStore(db *gorm.DB) *GormBackgroundStore {
	return &GormBackgroundStore{
		DB: db,
	}
}

func (s *GormBackgroundStore) IsValidName(name string) bool {
	err := s.DB.Where("name = ?", name).First(&models.Background{}).Error
	return err == nil
}
