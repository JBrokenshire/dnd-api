package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewRaceTrait(db *gorm.DB, rt *models.RaceTrait) {
	fillRaceTraitDetails(rt)
	db.Create(rt)
}

func fillRaceTraitDetails(rt *models.RaceTrait) {
	if rt.RaceID == 0 {
		rt.RaceID = 1
	}
	if rt.TraitID == 0 {
		rt.TraitID = 1
	}
}
