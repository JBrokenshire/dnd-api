package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewClass(db *gorm.DB, class *m.Class) {
	fillClassDefaults(class)
	err := db.Create(class).Error
	if err != nil {
		log.Println("Error creating class in factory: ", err.Error())
	}
}

func fillClassDefaults(class *m.Class) {
	if class.Name == "" {
		class.Name = random.String(16)
	}
	if class.ShortDescription == "" {
		class.ShortDescription = random.String(200)
	}
	if class.PrimaryAbility == "" {
		class.PrimaryAbility = "Strength"
	}
	if class.HitPointDieValue == 0 {
		class.HitPointDieValue = 10
	}
	if class.Saves == "" {
		class.Saves = "Strength & Constitution"
	}
}
