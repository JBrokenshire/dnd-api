package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetClassFeatures() {
	classFeatures := []models.ClassFeature{
		{
			ID:        1,
			ClassID:   3,
			FeatureID: 2,
		},
		{
			ID:        2,
			ClassID:   3,
			FeatureID: 3,
		},
		{
			ID:        3,
			ClassID:   3,
			FeatureID: 4,
		},
		{
			ID:        4,
			ClassID:   3,
			FeatureID: 5,
		},
		{
			ID:        5,
			ClassID:   3,
			FeatureID: 6,
			Level:     2,
		},
		{
			ID:        6,
			ClassID:   3,
			FeatureID: 7,
			Level:     2,
		},
		{
			ID:        7,
			ClassID:   3,
			FeatureID: 8,
			Level:     3,
		},
		{
			ID:        8,
			ClassID:   1,
			FeatureID: 2,
		},
		{
			ID:        9,
			ClassID:   1,
			FeatureID: 11,
		},
		{
			ID:        10,
			ClassID:   1,
			FeatureID: 12,
		},
		{
			ID:        11,
			ClassID:   1,
			FeatureID: 13,
			Level:     2,
		},
		{
			ID:        12,
			ClassID:   1,
			FeatureID: 14,
			Level:     2,
		},
		{
			ID:        13,
			ClassID:   1,
			FeatureID: 15,
			Level:     3,
		},
		{
			ID:        14,
			ClassID:   1,
			FeatureID: 1,
			Level:     4,
		},
		{
			ID:        15,
			ClassID:   1,
			FeatureID: 18,
			Level:     5,
		},
		{
			ID:        16,
			ClassID:   1,
			FeatureID: 19,
			Level:     5,
		},
		{
			ID:        17,
			ClassID:   6,
			FeatureID: 2,
		},
		{
			ID:        18,
			ClassID:   6,
			FeatureID: 6,
		},
		{
			ID:        19,
			ClassID:   6,
			FeatureID: 22,
		},
		{
			ID:        20,
			ClassID:   6,
			FeatureID: 23,
			Level:     2,
		},
		{
			ID:        21,
			ClassID:   6,
			FeatureID: 24,
			Level:     3,
		},
		{
			ID:        22,
			ClassID:   6,
			FeatureID: 1,
			Level:     4,
		},
		{
			ID:        23,
			ClassID:   6,
			FeatureID: 18,
			Level:     5,
		},
		{
			ID:        24,
			ClassID:   11,
			FeatureID: 2,
		},
		{
			ID:        25,
			ClassID:   11,
			FeatureID: 27,
		},
		{
			ID:        26,
			ClassID:   11,
			FeatureID: 28,
		},
		{
			ID:        27,
			ClassID:   11,
			FeatureID: 29,
			Level:     2,
		},
		{
			ID:        28,
			ClassID:   11,
			FeatureID: 30,
			Level:     2,
		},
		{
			ID:        29,
			ClassID:   11,
			FeatureID: 31,
			Level:     2,
		},
		{
			ID:        30,
			ClassID:   11,
			FeatureID: 32,
			Level:     3,
		},
		{
			ID:        24,
			ClassID:   11,
			FeatureID: 1,
			Level:     4,
		},
		{
			ID:        31,
			ClassID:   11,
			FeatureID: 33,
			Level:     5,
		},
	}

	for _, classFeature := range classFeatures {
		err := s.DB.Where("id = ?", classFeature.ID).FirstOrCreate(&classFeature).Error
		if err != nil {
			log.Printf("error creating class feature with id %v - %v", classFeature.ID, err)
		}
	}
}
