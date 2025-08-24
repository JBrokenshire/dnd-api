package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewCharacterSense(db *gorm.DB, sense *m.CharacterSense) {
	fillCharacterSenseDefaults(sense)
	err := db.Create(sense).Error
	if err != nil {
		log.Println("Error creating character sense in factory: ", err.Error())
	}
}

func fillCharacterSenseDefaults(sense *m.CharacterSense) {
	if sense.CharacterID == 0 {
		sense.CharacterID = 1
	}
	if sense.Sense == "" {
		sense.Sense = random.String(32)
	}
}
