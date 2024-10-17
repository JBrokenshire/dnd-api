package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterStore interface {
	Create(character *models.Character) error
	GetAll() ([]*models.Character, error)
	Get(id interface{}) (*models.Character, error)
	Update(character *models.Character) error
	Delete(id interface{}) error
}

type GormCharacterStore struct {
	DB *gorm.DB
}

func NewGormCharacterStore(db *gorm.DB) *GormCharacterStore {
	return &GormCharacterStore{
		DB: db,
	}
}

func (g *GormCharacterStore) Create(character *models.Character) error {
	return g.DB.Create(character).Error
}

func (g *GormCharacterStore) GetAll() ([]*models.Character, error) {
	var characters []*models.Character
	if err := g.DB.
		Preload("Class").
		Preload("Race").
		Find(&characters).
		Error; err != nil {
		return nil, err
	}
	return characters, nil
}

func (g *GormCharacterStore) Get(id interface{}) (*models.Character, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var character models.Character
	if err := g.DB.
		Preload("Class").
		Preload("Race").
		Preload("Background").
		Where("characters.id = ?", id).
		First(&character).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("character with id %q not found", id))
	}

	return &character, nil
}

func (g *GormCharacterStore) Update(character *models.Character) error {
	return g.DB.Save(&character).Error
}

func (g *GormCharacterStore) Delete(id interface{}) error {
	_, err := g.Get(id)
	if err != nil {
		return err
	}
	return g.DB.Delete(&models.Character{}, "id = ?", id).Error
}
