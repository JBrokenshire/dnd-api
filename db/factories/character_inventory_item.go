package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

func NewCharacterInventoryItem(db *gorm.DB, inventoryItem *m.CharacterInventoryItem) {
	fillCharacterInventoryItemDefaults(inventoryItem)
	err := db.Create(inventoryItem).Error
	if err != nil {
		log.Println("Error creating character inventory item in factory: ", err.Error())
	}
}

func fillCharacterInventoryItemDefaults(inventoryItem *m.CharacterInventoryItem) {
	if inventoryItem.CharacterId == 0 {
		inventoryItem.CharacterId = 1
	}
	if inventoryItem.ItemId == 0 {
		inventoryItem.ItemId = 1
	}
	if inventoryItem.Location == "" {
		inventoryItem.Location = "Equipment"
	}
	if inventoryItem.Quantity == 0 {
		inventoryItem.Quantity = 1
	}
}
