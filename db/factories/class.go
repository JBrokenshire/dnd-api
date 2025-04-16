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
}
