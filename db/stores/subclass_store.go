package stores

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

type SubclassStore interface {
	Get(id interface{}) (*models.Subclass, error)
	GetFeatures(id interface{}) ([]*models.Feature, error)
}

type GormSubclassStore struct {
	DB *gorm.DB
}

func NewGormSubclassStore(db *gorm.DB) *GormSubclassStore {
	return &GormSubclassStore{
		DB: db,
	}
}

func (s *GormSubclassStore) Get(id interface{}) (*models.Subclass, error) {
	var subclass models.Subclass
	err := s.DB.Where("id = ?", id).First(&subclass).Error
	if err != nil {
		return nil, err
	}
	return &subclass, nil
}

func (s *GormSubclassStore) GetFeatures(id interface{}) ([]*models.Feature, error) {
	var features []*models.Feature
	err := s.DB.
		Joins("JOIN subclass_features ON subclass_features.feature_id = features.id").
		Where("subclass_features.subclass_id = ?", id).
		Find(&features).Error
	if err != nil {
		return nil, err
	}
	return features, nil
}
