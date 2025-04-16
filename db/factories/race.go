package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewRace(db *gorm.DB, race *m.Race) {
	fillRaceDefaults(race)
	err := db.Create(race).Error
	if err != nil {
		log.Println("Error creating race in factory: ", err.Error())
	}
}

func fillRaceDefaults(race *m.Race) {
	if race.Name == "" {
		race.Name = random.String(16)
	}
}
