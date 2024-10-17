package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterSkillsAdvantageStore interface {
	GetSkillsAdvantagesByCharacterID(id interface{}) ([]*models.CharacterSkillAdvantage, error)
}

type GormCharacterSkillsAdvantagesStore struct {
	DB *gorm.DB
}

func NewGormCharacterSkillsAdvantagesStore(db *gorm.DB) *GormCharacterSkillsAdvantagesStore {
	return &GormCharacterSkillsAdvantagesStore{DB: db}
}

func (s *GormCharacterSkillsAdvantagesStore) GetSkillsAdvantagesByCharacterID(id interface{}) ([]*models.CharacterSkillAdvantage, error) {
	var characterSkillsAdvantages []*models.CharacterSkillAdvantage
	if err := s.DB.Table("character_skills_advantages").Where("character_id = ?", id).Find(&characterSkillsAdvantages).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("character skills advantages with character id: %q could not be found", id))
	}

	return characterSkillsAdvantages, nil
}
