package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewFeature(db *gorm.DB, f *models.Feature) {
	fillFeatureDetails(f)
	db.Create(f)
}

func fillFeatureDetails(f *models.Feature) {
	if f.Name == "" {
		f.Name = random.String(16)
	}
	if f.Description == "" {
		f.Description = random.String(64)
	}
}
