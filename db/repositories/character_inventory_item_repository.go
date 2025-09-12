package repositories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

type CharacterInventoryItemRepository struct {
	*Repository
}

func NewCharacterInventoryItemRepository(db *gorm.DB) *CharacterInventoryItemRepository {
	return &CharacterInventoryItemRepository{
		&Repository{
			Db: db,
		},
	}
}

func (r *CharacterInventoryItemRepository) GetById(characterID, inventoryItemId interface{}) *m.CharacterInventoryItem {
	var characterInventoryItem m.CharacterInventoryItem
	r.Db.
		Preload("Item").
		Where("id = ?", inventoryItemId).
		Where("character_id = ?", characterID).
		First(&characterInventoryItem)
	return &characterInventoryItem
}
