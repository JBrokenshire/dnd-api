package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

func NewCharacterProficientSkill(db *gorm.DB, characterProficientSkill *m.CharacterProficientSkill) {
	fillCharacterProficientSkillDefaults(characterProficientSkill)
	err := db.Create(characterProficientSkill).Error
	if err != nil {
		log.Println("Error creating character proficient skill in factory: ", err.Error())
	}
}

func fillCharacterProficientSkillDefaults(profSkill *m.CharacterProficientSkill) {
	if profSkill.Skill == "" {
		profSkill.Skill = m.ValidSkills[0]
	}
	if profSkill.ProficiencyType == "" {
		profSkill.ProficiencyType = m.ValidProficiencyTypes[0]
	}
}
