package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) Create(object interface{}) error {
	return r.Db.Create(object).Error
}

func (r *Repository) Update(object interface{}) error {
	return r.Db.Save(object).Error
}

func (r *Repository) Delete(object interface{}) error {
	return r.Db.Delete(object).Error
}
