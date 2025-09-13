package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewTrait(db *gorm.DB, trait *m.Trait) {
	fillTraitDefaults(trait)
	err := db.Create(trait).Error
	if err != nil {
		log.Println("Error creating trait in factory: ", err.Error())
	}
}

func fillTraitDefaults(trait *m.Trait) {
	if trait.ID == "" {
		trait.ID = random.String(16)
	}
	if trait.Name == "" {
		trait.Name = random.String(16)
	}
	if trait.Description == "" {
		trait.Description = random.String(16)
	}
}
