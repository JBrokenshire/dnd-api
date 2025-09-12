package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

func NewArmour(db *gorm.DB, armour *m.Armour) {
	fillArmourDefaults(armour)
	err := db.Create(armour).Error
	if err != nil {
		log.Println("Error creating armour in factory: ", err.Error())
	}
}

func fillArmourDefaults(armour *m.Armour) {
	if armour.BaseAC == 0 {
		armour.BaseAC = 12
	}
	if armour.Type == "" {
		armour.Type = m.ArmourTypeLight
	}
}
