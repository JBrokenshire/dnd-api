package repositories

import (
	"dnd-api/db"
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type RaceRepository struct {
	*Repository
}

func NewRaceRepository(db *gorm.DB) *RaceRepository {
	return &RaceRepository{
		Repository: &Repository{
			Db: db,
		},
	}
}

func (r *RaceRepository) GetRaces(c echo.Context, scopes Scopes) ([]*m.Race, int, int) {
	var races []*m.Race
	page, pageSize, paginateFunc := db.Paginate(c)
	r.Db.
		Scopes(paginateFunc).
		Scopes(scopes...).
		Find(&races)
	return races, page, pageSize
}

func (r *RaceRepository) CountRaces(scopes Scopes) int {
	var count int64
	r.Db.
		Model(&m.Race{}).
		Scopes(scopes...).
		Count(&count)
	return int(count)
}

func (r *RaceRepository) GetById(id interface{}) *m.Race {
	var race m.Race
	r.Db.Where("id = ?", id).First(&race)
	return &race
}
