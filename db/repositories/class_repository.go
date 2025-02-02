package repositories

import (
	"github.com/JBrokenshire/dnd-api/db"
	m "github.com/JBrokenshire/dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ClassRepository struct {
	*Repository
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{
		Repository: &Repository{
			Db: db,
		},
	}
}

func (r *ClassRepository) GetClasses(c echo.Context, scopes []func(db *gorm.DB) *gorm.DB) ([]*m.Class, int, int) {
	var classes []*m.Class
	page, pageSize, paginateFunc := db.Paginate(c)
	r.Db.
		Scopes(paginateFunc).
		Scopes(scopes...).
		Find(&classes)
	return classes, page, pageSize
}

func (r *ClassRepository) Count(scopes []func(db *gorm.DB) *gorm.DB) int {
	var count int64
	r.Db.
		Model(&m.Class{}).
		Scopes(scopes...).
		Count(&count)
	return int(count)
}
