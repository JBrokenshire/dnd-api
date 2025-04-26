package repositories

import (
	"dnd-api/db"
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type SubclassRepository struct {
	*Repository
}

func NewSubclassRepository(db *gorm.DB) *SubclassRepository {
	return &SubclassRepository{
		&Repository{
			Db: db,
		},
	}
}

func (r *SubclassRepository) GetSubclasses(c echo.Context, classId interface{}, scopes Scopes) ([]*m.Subclass, int, int) {
	var subclasses []*m.Subclass
	page, pageSize, paginateFunc := db.Paginate(c)
	r.Db.
		Scopes(paginateFunc).
		Scopes(scopes...).
		Where("class_id = ?", classId).
		Order("name ASC").
		Find(&subclasses)

	for i := range subclasses {
		r.Db.Where("model = ?", m.FileModelSubclassLogo).Where("model_id = ?", subclasses[i].ID).Take(&subclasses[i].Logo)
	}

	return subclasses, page, pageSize
}

func (r *SubclassRepository) CountSubclasses(classId interface{}, scopes Scopes) int {
	var count int64
	r.Db.
		Model(&m.Subclass{}).
		Scopes(scopes...).
		Where("class_id = ?", classId).
		Count(&count)
	return int(count)
}

func (r *SubclassRepository) GetById(subclassId, classId interface{}) *m.Subclass {
	var subclass m.Subclass
	r.Db.
		Where("id = ?", subclassId).
		Where("class_id = ?", classId).
		First(&subclass)

	// Load class logo
	r.Db.Where("model = ?", m.FileModelSubclassLogo).Where("model_id = ?", subclass.ID).Take(&subclass.Logo)

	return &subclass
}
