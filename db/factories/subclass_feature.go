package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewSubclassFeature(db *gorm.DB, sf *models.SubclassFeature) {
	fillSubclassFeatureDetails(sf)
	db.Create(sf)
}

func fillSubclassFeatureDetails(sf *models.SubclassFeature) {
	if sf.SubclassID == 0 {
		sf.SubclassID = 1
	}
	if sf.FeatureID == 0 {
		sf.FeatureID = 1
	}
}
