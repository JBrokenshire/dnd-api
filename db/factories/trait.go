package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewTrait(db *gorm.DB, t *models.Trait) {
	fillTraitDetails(t)
	db.Create(t)
}

func fillTraitDetails(t *models.Trait) {
	if t.Name == "" {
		t.Name = random.String(16)
	}
}
