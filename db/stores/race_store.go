package stores

import (
	"dnd-api/db/models"
	"errors"
	"github.com/jinzhu/gorm"
	"reflect"
)

type RaceStore interface {
	GetAll() ([]*models.Race, error)
	Get(id interface{}) (*models.Race, error)
	GetTraits(id interface{}) ([]*models.Trait, error)
	IsValidID(id interface{}) bool
}

type GormRaceStore struct {
	DB *gorm.DB
}

func NewGormRaceStore(db *gorm.DB) *GormRaceStore {
	return &GormRaceStore{
		DB: db,
	}
}

func (s *GormRaceStore) GetAll() ([]*models.Race, error) {
	var races []*models.Race
	if err := s.DB.Find(&races).Error; err != nil {
		return nil, err
	}

	return races, nil
}

func (s *GormRaceStore) Get(id interface{}) (*models.Race, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var race models.Race
	if err := s.DB.Where("id = ?", id).First(&race).Error; err != nil {
		return nil, err
	}
	return &race, nil
}

func (s *GormRaceStore) GetTraits(id interface{}) ([]*models.Trait, error) {
	var traits []*models.Trait
	err := s.DB.
		Joins("JOIN race_traits ON race_traits.trait_id = traits.id").
		Where("race_traits.race_id = ?", id).
		Find(&traits).Error
	if err != nil {
		return nil, err
	}

	return traits, nil
}

func (s *GormRaceStore) IsValidID(id interface{}) bool {
	var race models.Race
	if err := s.DB.Where("id = ?", id).First(&race).Error; err != nil {
		return false
	}
	return true
}
