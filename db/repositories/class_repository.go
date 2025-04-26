package repositories

import (
	"dnd-api/db"
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ClassRepository struct {
	*Repository
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{
		&Repository{
			Db: db,
		},
	}
}

func (r *ClassRepository) GetClasses(c echo.Context, scopes Scopes) ([]*m.Class, int, int) {
	var classes []*m.Class
	page, pageSize, paginateFunc := db.Paginate(c)
	r.Db.
		Scopes(paginateFunc).
		Scopes(scopes...).
		Order("name ASC").
		Find(&classes)

	// Load on images
	for i := range classes {
		r.Db.Where("model = ?", m.FileModelClassLogo).Where("model_id = ?", classes[i].ID).Take(&classes[i].Logo)
	}

	return classes, page, pageSize
}

func (r *ClassRepository) CountClasses(scopes Scopes) int {
	var count int64
	r.Db.
		Model(&m.Class{}).
		Scopes(scopes...).
		Count(&count)
	return int(count)
}

func (r *ClassRepository) GetById(id interface{}) *m.Class {
	var class m.Class
	r.Db.Where("id = ?", id).First(&class)

	// Load class logo
	r.Db.Where("model = ?", m.FileModelClassLogo).Where("model_id = ?", class.ID).Take(&class.Logo)

	return &class
}

func (r *ClassRepository) GetSubclasses(id interface{}) []*m.Subclass {
	var subclasses []*m.Subclass
	r.Db.
		Where("class_id = ?", id).
		Find(&subclasses)

	return subclasses
}
