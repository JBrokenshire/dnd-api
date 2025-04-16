package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewCharacter(db *gorm.DB, character *m.Character) {
	fillCharacterDefaults(character)
	err := db.Create(character).Error
	if err != nil {
		log.Println("Error creating character in factory: ", err.Error())
	}
}

func fillCharacterDefaults(character *m.Character) {
	if character.UserId == 0 {
		character.UserId = 1
	}
	if character.Name == "" {
		character.Name = random.String(16)
	}
}
