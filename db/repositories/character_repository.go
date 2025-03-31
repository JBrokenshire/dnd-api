package repositories

import (
	"dnd-api/db"
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type CharacterRepository struct {
	*Repository
}

func NewCharacterRepository(db *gorm.DB) *CharacterRepository {
	return &CharacterRepository{
		&Repository{Db: db},
	}
}

func (r *CharacterRepository) GetCharacters(c echo.Context) ([]m.Character, int, int) {
	var characters []m.Character
	page, pageSize, paginateFunc := db.Paginate(c)
	r.Db.
		Preload("Race").
		Preload("Class").
		Scopes(paginateFunc).
		Find(&characters)

	return characters, page, pageSize
}

func (r *CharacterRepository) Count() int {
	var count int
	r.Db.Model(&m.Character{}).Count(&count)
	return count
}

func (r *CharacterRepository) GetCharacter(id interface{}) *m.Character {
	var character m.Character
	r.Db.
		Preload("Race").
		Preload("Class").
		Where("id = ?", id).
		First(&character)
	return &character
}
