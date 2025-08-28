package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

func NewCharacterDefense(db *gorm.DB, characterDefense *m.CharacterDefense) {
	fillCharacterDefenseDefaults(characterDefense)
	err := db.Create(characterDefense).Error
	if err != nil {
		log.Println("Error creating character defense in factory: ", err.Error())
	}
}

func fillCharacterDefenseDefaults(defense *m.CharacterDefense) {
	if defense.DamageType == "" {
		defense.DamageType = m.DamageTypeFire
	}
	if defense.DefenseType == "" {
		defense.DefenseType = m.DefenseTypeResistance
	}
}
