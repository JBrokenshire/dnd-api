package stores

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

type ArmourStore interface {
	Get(id interface{}) (models.Armour, error)
}

type GormArmourStore struct {
	DB *gorm.DB
}

func NewGormArmourStore(db *gorm.DB) *GormArmourStore {
	return &GormArmourStore{
		DB: db,
	}
}

func (s *GormArmourStore) Get(id interface{}) (models.Armour, error) {
	var armour models.Armour
	if err := s.DB.Table("armour").Where("item_id = ?", id).Find(&armour).Error; err != nil {
		return models.Armour{}, err
	}
	return armour, nil
}
