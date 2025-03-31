package factories

import (
	m "dnd-api/db/models"
	"dnd-api/pkg/random"
	"github.com/jinzhu/gorm"
	"log"
)

func NewRace(db *gorm.DB, race *m.Race) {
	fillRaceDefaults(race)
	err := db.Create(race).Error
	if err != nil {
		log.Printf("Error creating race in factory: %v\n", err.Error())
	}
}

func fillRaceDefaults(race *m.Race) {
	if race.Name == "" {
		race.Name = random.String()
	}
}
