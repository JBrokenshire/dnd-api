package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CharacterSkillsAdvantageStore interface {
	Update(characterSkillAdvantage *models.CharacterSkillAdvantage) error
	GetSkillsAdvantagesByCharacterID(id interface{}) ([]*models.CharacterSkillAdvantage, error)
	GetByCharacterIDAndSkillName(id interface{}, skillName string) (*models.CharacterSkillAdvantage, error)
}

type GormCharacterSkillsAdvantagesStore struct {
	DB *gorm.DB
}

func NewGormCharacterSkillsAdvantagesStore(db *gorm.DB) *GormCharacterSkillsAdvantagesStore {
	return &GormCharacterSkillsAdvantagesStore{DB: db}
}

func (s *GormCharacterSkillsAdvantagesStore) GetSkillsAdvantagesByCharacterID(id interface{}) ([]*models.CharacterSkillAdvantage, error) {
	var characterSkillsAdvantages []*models.CharacterSkillAdvantage
	if err := s.DB.
		Table("character_skills_advantages").
		Where("character_id = ?", id).
		Find(&characterSkillsAdvantages).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("character skills advantages with character id: %q could not be found", id))
	}

	return characterSkillsAdvantages, nil
}

func (s *GormCharacterSkillsAdvantagesStore) GetByCharacterIDAndSkillName(id interface{}, skillName string) (*models.CharacterSkillAdvantage, error) {
	var characterSkillsAdvantage models.CharacterSkillAdvantage
	if err := s.DB.
		Table("character_skills_advantages").
		Where("character_id = ? AND skill_name = ?", id, skillName).
		First(&characterSkillsAdvantage).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("character skills advantages with character id: %q could not be found", id))
	}

	return &characterSkillsAdvantage, nil
}

func (s *GormCharacterSkillsAdvantagesStore) Update(characterSkillAdvantage *models.CharacterSkillAdvantage) error {
	return s.DB.Table("character_skills_advantages").Save(&characterSkillAdvantage).Error
}
