package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewSubclass(db *gorm.DB, subclass *models.Subclass) {
	fillSubclassDefaults(subclass)
	err := db.Create(subclass).Error
	if err != nil {
		log.Println("Error creating subclass in factory: ", err.Error())
	}
}

func fillSubclassDefaults(subclass *models.Subclass) {
	if subclass.ClassId == 0 {
		subclass.ClassId = 1
	}
	if subclass.Name == "" {
		subclass.Name = random.String(16)
	}
	if subclass.ShortDescription == "" {
		subclass.ShortDescription = random.String(32)
	}
}
