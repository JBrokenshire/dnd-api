package stores

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

type CharacterFeaturesStore interface {
	GetClassFeatures(classID interface{}, level int) ([]*models.ClassFeature, error)
	GetSubclassFeatures(subclassID interface{}, level int) ([]*models.SubclassFeature, error)
}

type GormCharacterFeaturesStore struct {
	DB *gorm.DB
}

func NewGormCharacterFeaturesStore(db *gorm.DB) *GormCharacterFeaturesStore {
	return &GormCharacterFeaturesStore{
		DB: db,
	}
}

func (s *GormCharacterFeaturesStore) GetClassFeatures(classID interface{}, level int) ([]*models.ClassFeature, error) {
	var features []*models.ClassFeature
	err := s.DB.
		Preload("Feature").
		Where("class_id = ?", classID).
		Where("class_features.level <= ?", level).
		Order("class_features.level ASC").
		Find(&features).Error
	if err != nil {
		return nil, err
	}

	return features, nil
}

func (s *GormCharacterFeaturesStore) GetSubclassFeatures(subclassID interface{}, level int) ([]*models.SubclassFeature, error) {
	var features []*models.SubclassFeature
	err := s.DB.
		Preload("Feature").
		Where("subclass_id = ?", subclassID).
		Where("subclass_features.level <= ?", level).
		Order("subclass_features.level ASC").
		Find(&features).Error
	if err != nil {
		return nil, err
	}

	return features, nil
}
