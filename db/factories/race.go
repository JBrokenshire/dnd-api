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
	if race.ShortDescription == "" {
		race.ShortDescription = random.String(32)
	}
	if race.CreatureType == "" {
		race.CreatureType = "Humanoid"
	}
	if race.Size == "" {
		race.Size = "Medium"
	}
	if race.BaseSpeed == 0 {
		race.BaseSpeed = 30
	}
}
