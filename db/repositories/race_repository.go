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

func (r *RaceRepository) GetRaces(c echo.Context) ([]m.Race, int, int) {
	var races []m.Race
	page, pageSize, paginateFunc := db.Paginate(c)
	r.Db.Scopes(paginateFunc).Find(&races)
	return races, page, pageSize
}

func (r *RaceRepository) Count() int {
	var count int
	r.Db.Model(&m.Race{}).Count(&count)
	return count
}

func (r *RaceRepository) GetRace(id interface{}) *m.Race {
	var race m.Race
	r.Db.Where("id = ?", id).First(&race)
	return &race
}
