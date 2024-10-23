package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetSubclassFeatures() {
	subclassFeatures := []models.SubclassFeature{
		{
			ID:         1,
			SubclassID: 2,
			FeatureID:  9,
			Level:      3,
		},
		{
			ID:         2,
			SubclassID: 2,
			FeatureID:  10,
			Level:      3,
		},
		{
			ID:         3,
			SubclassID: 1,
			FeatureID:  16,
			Level:      3,
		},
		{
			ID:         4,
			SubclassID: 1,
			FeatureID:  17,
			Level:      3,
		},
		{
			ID:         5,
			SubclassID: 3,
			FeatureID:  20,
			Level:      3,
		},
		{
			ID:         6,
			SubclassID: 3,
			FeatureID:  21,
			Level:      3,
		},
		{
			ID:         7,
			SubclassID: 4,
			FeatureID:  25,
			Level:      3,
		},
		{
			ID:         8,
			SubclassID: 4,
			FeatureID:  26,
			Level:      3,
		},
	}

	for _, subclassFeature := range subclassFeatures {
		err := s.DB.Where("id = ?", subclassFeature.ID).FirstOrCreate(&subclassFeature).Error
		if err != nil {
			log.Printf("error creating subclass feature with id %v - %v", subclassFeature.ID, err)
		}
	}
}
