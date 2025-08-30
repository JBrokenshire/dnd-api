package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

func NewCharacterSpell(db *gorm.DB, cs *m.CharacterSpell) {
	err := db.Create(cs).Error
	if err != nil {
		log.Println("Error creating character_spell in factory: ", err.Error())
	}
}
