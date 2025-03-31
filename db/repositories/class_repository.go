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

func (r *ClassRepository) GetClasses(c echo.Context) ([]m.Class, int, int) {
	var classes []m.Class
	page, pageSize, paginateFunc := db.Paginate(c)
	r.Db.Scopes(paginateFunc).Find(&classes)
	return classes, page, pageSize
}

func (r *ClassRepository) Count() int {
	var count int
	r.Db.Model(&m.Class{}).Count(&count)
	return count
}

func (r *ClassRepository) GetClass(id interface{}) *m.Class {
	var class m.Class
	r.Db.Where("id = ?", id).First(&class)
	return &class
}
