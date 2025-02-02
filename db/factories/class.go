package factories

import (
	m "github.com/JBrokenshire/dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewClass(db *gorm.DB, class *m.Class) {
	fillClassDetails(class)
	err := db.Create(class).Error
	if err != nil {
		log.Printf("Error creating class: %v", err)
	}
}

func fillClassDetails(class *m.Class) {
	if class.Name == "" {
		class.Name = random.String(16)
	}
	if class.ShortDescription == "" {
		class.ShortDescription = random.String(32)
	}
	if class.LongDescription == "" {
		class.LongDescription = random.String(64)
	}
}
