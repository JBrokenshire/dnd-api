package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewItem(db *gorm.DB, item *m.Item) {
	fillItemDefaults(item)
	err := db.Create(item).Error
	if err != nil {
		log.Println("Error creating item in factory: ", err.Error())
	}
}

func fillItemDefaults(item *m.Item) {
	if item.Name == "" {
		item.Name = random.String(16)
	}
	if item.Rarity == "" {
		item.Rarity = m.RarityCommon
	}
	if item.Type == "" {
		item.Type = m.ItemTypeItem
	}
}
