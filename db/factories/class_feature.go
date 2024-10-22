package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewClassFeature(db *gorm.DB, cf *models.ClassFeature) {
	fillClassFeatureDetails(cf)
	db.Create(cf)
}

func fillClassFeatureDetails(cf *models.ClassFeature) {
	if cf.ClassID == 0 {
		cf.ClassID = 1
	}
	if cf.FeatureID == 0 {
		cf.FeatureID = 1
	}
}
