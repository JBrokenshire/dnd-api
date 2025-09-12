package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

func NewClassSpellLevel(db *gorm.DB, classSpellLevel *m.ClassSpellLevel) {
	fillClassSpellLevelDefaults(classSpellLevel)
	err := db.Create(classSpellLevel).Error
	if err != nil {
		log.Println("Error creating class spell level in factory: ", err.Error())
	}
}

func fillClassSpellLevelDefaults(classSpellLevel *m.ClassSpellLevel) {
	if classSpellLevel.ClassLevel == 0 {
		classSpellLevel.ClassLevel = 1
	}
	if classSpellLevel.SpellLevel == 0 {
		classSpellLevel.SpellLevel = 1
	}
	if classSpellLevel.NumberOfSlots == 0 {
		classSpellLevel.NumberOfSlots = 1
	}
}
