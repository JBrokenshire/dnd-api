package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

func NewRaceTrait(db *gorm.DB, raceTrait *m.RaceTrait) {
	err := db.Create(raceTrait).Error
	if err != nil {
		log.Println("Error creating race trait in factory: ", err.Error())
	}
}
