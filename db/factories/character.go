package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacter(db *gorm.DB, character *models.Character) {
	fillCharacterDetails(character)
	db.Create(character)
}

func fillCharacterDetails(c *models.Character) {
	if c.Name == "" {
		c.Name = "TEST CHARACTER"
	}
	if c.Level == 0 {
		c.Level = 1
	}
	if c.ClassID == 0 {
		c.ClassID = 1
	}
	if c.RaceID == 0 {
		c.RaceID = 1
	}
	if c.Strength == 0 {
		c.Strength = 10
	}
	if c.Dexterity == 0 {
		c.Dexterity = 10
	}
	if c.Constitution == 0 {
		c.Constitution = 10
	}
	if c.Intelligence == 0 {
		c.Intelligence = 10
	}
	if c.Wisdom == 0 {
		c.Wisdom = 10
	}
	if c.Charisma == 0 {
		c.Charisma = 10
	}
	if c.CurrentHitPoints == 0 {
		c.CurrentHitPoints = 1
	}
	if c.MaxHitPoints == 0 {
		c.MaxHitPoints = 1
	}
	if c.AttacksPerAction == 0 {
		c.AttacksPerAction = 1
	}
	if c.BackgroundName == "" {
		c.BackgroundName = "Outlander"
	}
}
