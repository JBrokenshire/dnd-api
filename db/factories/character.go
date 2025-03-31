package factories

import (
	m "dnd-api/db/models"
	"dnd-api/pkg/random"
	"github.com/jinzhu/gorm"
	"log"
)

func NewCharacter(db *gorm.DB, character *m.Character) {
	fillCharacterDefaults(character)
	err := db.Create(character).Error
	if err != nil {
		log.Printf("Error creating character in factory: %v\n", err.Error())
	}
}

func fillCharacterDefaults(character *m.Character) {
	if character.Name == "" {
		character.Name = random.String()
	}
	if character.RaceId == 0 {
		character.RaceId = 1
	}
	if character.ClassId == 0 {
		character.ClassId = 1
	}
}
