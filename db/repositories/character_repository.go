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
		&Repository{
			Db: db,
		},
	}
}

func (r *CharacterRepository) GetCharacters(c echo.Context, userId interface{}, scopes []func(db *gorm.DB) *gorm.DB) ([]*m.Character, int, int) {
	var characters []*m.Character
	page, pageSize, paginateFunc := db.Paginate(c)
	r.Db.
		Preload("Class").
		Preload("Race").
		Scopes(paginateFunc).
		Scopes(scopes...).
		Where("user_id = ?", userId).
		Find(&characters)

	return characters, page, pageSize
}

func (r *CharacterRepository) CountCharacters(userId interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	var count int64
	r.Db.
		Model(&m.Character{}).
		Scopes(scopes...).
		Where("user_id = ?", userId).
		Count(&count)
	return int(count)
}

func (r *CharacterRepository) GetById(id interface{}, userId interface{}) *m.Character {
	var character m.Character
	r.Db.
		Preload("Class").
		Preload("Race").
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Find(&character)
	return &character
}
