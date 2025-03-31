package factories

import (
	"dnd-api/db/models"
	"dnd-api/pkg/random"
	"github.com/jinzhu/gorm"
	"log"
)

func NewClass(db *gorm.DB, class *models.Class) {
	fillClassDefaults(class)
	err := db.Create(class).Error
	if err != nil {
		log.Printf("Error creating class in factory: %v\n", err.Error())
	}
}

func fillClassDefaults(class *models.Class) {
	if class.Name == "" {
		class.Name = random.String()
	}
}
