package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewRace(db *gorm.DB, r *models.Race) {
	fillRaceDetails(r)
	db.Create(r)
}

func fillRaceDetails(r *models.Race) {
	if r.Name == "" {
		r.Name = random.String(16)
	}
	if r.ShortDescription == "" {
		r.ShortDescription = random.String(16)
	}
	if r.LongDescription == "" {
		r.LongDescription = random.String(32)
	}
	if r.BaseWalkingSpeed == 0 {
		r.BaseWalkingSpeed = 30
	}
}
