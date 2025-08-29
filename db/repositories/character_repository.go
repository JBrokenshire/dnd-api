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

	// Load on images
	for i := range characters {
		r.Db.Where("model = ?", m.FileModelCharacterProfilePicture).Where("model_id = ?", characters[i].ID).Take(&characters[i].ProfilePicture)
	}

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

	// Load images
	r.Db.Where("model = ?", m.FileModelCharacterProfilePicture).Where("model_id = ?", character.ID).Take(&character.ProfilePicture)
	r.Db.Where("model = ?", m.FileModelClassBackgroundImage).Where("model_id = ?", character.Class.ID).Take(&character.Class.BackgroundImage)

	// Load Proficient Skills
	r.Db.Where("character_id = ?", character.ID).Find(&character.ProficientSkills)
	// Load Defenses
	r.Db.Where("character_id = ?", character.ID).Find(&character.Defenses)
	// Load Background
	if character.BackgroundId != 0 {
		r.Db.Where("id = ?", character.BackgroundId).Find(&character.Background)
	}

	return &character
}
