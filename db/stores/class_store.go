package stores

import (
	"dnd-api/db/models"
	"errors"
	"github.com/jinzhu/gorm"
	"reflect"
)

type ClassStore interface {
	GetAll() ([]*models.Class, error)
	Get(id interface{}) (*models.Class, error)
	Update(class *models.Class) error
	IsValidID(id interface{}) bool
}

type GormClassStore struct {
	db *gorm.DB
}

func NewGormClassStore(db *gorm.DB) *GormClassStore {
	return &GormClassStore{db: db}
}

func (s *GormClassStore) GetAll() ([]*models.Class, error) {
	var classes []*models.Class
	if err := s.db.Find(&classes).Error; err != nil {
		return nil, err
	}

	return classes, nil
}

func (s *GormClassStore) Get(id interface{}) (*models.Class, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var class models.Class
	if err := s.db.Where("id = ?", id).First(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (s *GormClassStore) Update(class *models.Class) error {
	return s.db.Save(class).Error
}

func (s *GormClassStore) IsValidID(id interface{}) bool {
	var class models.Class
	if err := s.db.Where("id = ?", id).First(&class).Error; err != nil {
		return false
	}
	return true
}
