package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterInventoryStore interface {
	GetInventoryByCharacterID(id interface{}) ([]*models.CharacterInventoryItem, error)
	GetEquippedWeaponsByCharacterID(id interface{}) ([]*models.Weapon, error)
	GetEquippedArmourByCharacterID(id interface{}) (*models.Armour, error)
	GetEquippedShieldByCharacterID(id interface{}) (*models.Shield, error)
	GetCharacterInventoryItemByID(characterID interface{}, itemID interface{}) (*models.CharacterInventoryItem, error)
	UpdateCharacterInventoryItem(characterInventoryItem *models.CharacterInventoryItem) error
}

type GormCharacterInventoryStore struct {
	DB *gorm.DB
}

func NewGormCharacterInventoryStore(db *gorm.DB) *GormCharacterInventoryStore {
	return &GormCharacterInventoryStore{
		DB: db,
	}
}

func (g *GormCharacterInventoryStore) GetInventoryByCharacterID(id interface{}) ([]*models.CharacterInventoryItem, error) {
	var characterInventory []*models.CharacterInventoryItem
	err := g.DB.
		Preload("Item").
		Where("character_id = ?", id).
		Find(&characterInventory).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("inventory items with character id: %q could not be found", id))
	}

	return characterInventory, nil
}

func (g *GormCharacterInventoryStore) GetEquippedWeaponsByCharacterID(id interface{}) ([]*models.Weapon, error) {
	var characterEquippedWeapons []*models.CharacterInventoryItem
	err := g.DB.
		Table("character_inventory_items").
		Where("character_id = ? AND location = 'Equipment' AND equipped = true AND type = 'weapon'", id).
		Find(&characterEquippedWeapons).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("inventory items with character id: %q could not be found", id))
	}

	var weapons []*models.Weapon
	for _, equippedWeapon := range characterEquippedWeapons {
		var weapon models.Weapon
		_ = g.DB.
			Preload("Item").
			Where("item_id = ?", equippedWeapon.ItemID).
			First(&weapon).Error
		if weapon.ItemID != 0 {
			weapons = append(weapons, &weapon)
		}
	}

	return weapons, nil
}

func (g *GormCharacterInventoryStore) GetEquippedArmourByCharacterID(id interface{}) (*models.Armour, error) {
	var characterEquippedArmour models.CharacterInventoryItem
	_ = g.DB.
		Table("character_inventory_items").
		Where("character_id = ? AND location = 'Equipment' AND equipped = true AND type = 'armour'", id).
		First(&characterEquippedArmour).Error
	if characterEquippedArmour.ID == 0 {
		return nil, nil
	}

	var armour models.Armour
	err := g.DB.
		Table("armour").
		Preload("Item").
		Where("item_id = ?", characterEquippedArmour.ItemID).
		First(&armour).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("armour with item id: %q could not be found", characterEquippedArmour.ItemID))
	}

	return &armour, nil
}

func (g *GormCharacterInventoryStore) GetEquippedShieldByCharacterID(id interface{}) (*models.Shield, error) {
	var characterEquippedShield models.CharacterInventoryItem
	_ = g.DB.
		Table("character_inventory_items").
		Where("character_id = ? AND location = 'Equipment' AND equipped = true AND type = 'shield'", id).
		First(&characterEquippedShield).Error
	if characterEquippedShield.ID == 0 {
		return nil, nil
	}

	var shield models.Shield
	err := g.DB.
		Preload("Item").
		Where("item_id = ?", characterEquippedShield.ItemID).
		First(&shield).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("shield with item id: %q could not be found", id))
	}

	return &shield, nil
}

func (g *GormCharacterInventoryStore) GetCharacterInventoryItemByID(characterID interface{}, itemID interface{}) (*models.CharacterInventoryItem, error) {
	var character models.Character
	err := g.DB.Table("characters").Where("id = ?", characterID).Find(&character).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error getting character with id %v: %v", characterID, err))
	}

	var characterInventoryItem models.CharacterInventoryItem
	err = g.DB.
		Preload("Item").
		Where("id = ?", itemID).
		First(&characterInventoryItem).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error getting character inventory item with id %v: %v", itemID, err))
	}

	return &characterInventoryItem, nil
}

func (g *GormCharacterInventoryStore) UpdateCharacterInventoryItem(item *models.CharacterInventoryItem) error {
	return g.DB.Save(&item).Error
}
