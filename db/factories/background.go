package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewBackground(db *gorm.DB, background *m.Background) {
	fillBackgroundDefaults(background)
	err := db.Create(background).Error
	if err != nil {
		log.Println("Error creating background in factory: ", err.Error())
	}
}

func fillBackgroundDefaults(b *m.Background) {
	if b.Name == "" {
		b.Name = random.String(32)
	}
	if b.Feature == "" {
		b.Feature = random.String(32)
	}
	if b.Description == "" {
		b.Description = random.String(32)
	}
}
