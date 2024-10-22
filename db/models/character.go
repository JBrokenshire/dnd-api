package models

import (
	"dnd-api/utils"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Character struct {
	ID                int    `gorm:"autoIncrement;primary_key" json:"id"`
	Name              string `gorm:"not null" json:"name"`
	Level             int    `gorm:"not null" json:"level"`
	ProfilePictureURL string `json:"profile_picture_url"`
	ClassID           int    `json:"class_id"`
	RaceID            int    `json:"race_id"`
	SubclassID        *int   `json:"subclass_id"`

	Strength               int  `gorm:"not null" json:"strength"`
	Dexterity              int  `gorm:"not null" json:"dexterity"`
	Constitution           int  `gorm:"not null" json:"constitution"`
	Intelligence           int  `gorm:"not null" json:"intelligence"`
	Wisdom                 int  `gorm:"not null" json:"wisdom"`
	Charisma               int  `gorm:"not null" json:"charisma"`
	ProficientStrength     bool `gorm:"not null" json:"proficient_strength"`
	ProficientDexterity    bool `gorm:"not null" json:"proficient_dexterity"`
	ProficientConstitution bool `gorm:"not null" json:"proficient_constitution"`
	ProficientIntelligence bool `gorm:"not null" json:"proficient_intelligence"`
	ProficientWisdom       bool `gorm:"not null" json:"proficient_wisdom"`
	ProficientCharisma     bool `gorm:"not null" json:"proficient_charisma"`

	WalkingSpeedModifier int  `gorm:"not null" json:"walking_speed_modifier"`
	Inspiration          bool `gorm:"not null" json:"inspiration"`
	CurrentHitPoints     int  `gorm:"not null" json:"current_hit_points"`
	MaxHitPoints         int  `gorm:"not null" json:"max_hit_points"`
	TempHitPoints        int  `gorm:"not null" json:"temp_hit_points"`

	InitiativeModifier int `gorm:"not null" json:"initiative_modifier"`
	AttacksPerAction   int `gorm:"not null" json:"attacks_per_action"`

	BackgroundName string `gorm:"not null" json:"background_name"`
	Alignment      string `json:"alignment"`
	Gender         string `json:"gender"`
	Eyes           string `json:"eyes"`
	Size           string `json:"size"`
	Height         string `json:"height"`
	Faith          string `json:"faith"`
	Hair           string `json:"hair"`
	Skin           string `json:"skin"`
	Age            int    `json:"age"`
	Weight         int    `json:"weight"`

	Organisations string `json:"organisations"`
	Allies        string `json:"allies"`
	Enemies       string `json:"enemies"`
	Backstory     string `json:"backstory"`

	Class      Class      `json:"class"`
	Race       Race       `json:"race"`
	Subclass   *Subclass  `json:"subclass"`
	Background Background `json:"background"`
}

func (c *Character) BeforeCreate(db *gorm.DB) error {
	var class Class
	err := db.Where("id = ?", c.ClassID).Find(&class).Error
	if err != nil {
		return fmt.Errorf("class with id '%v' not found - %v", c.ClassID, err)
	}

	var race Race
	err = db.Where("id = ?", c.RaceID).Find(&race).Error
	if err != nil {
		return fmt.Errorf("race with id '%v' not found - %v", c.RaceID, err)
	}

	if c.Subclass != nil {
		var subclass Subclass
		err = db.Where("id = ?", c.Subclass.ID).Find(&subclass).Error
		if err != nil {
			return fmt.Errorf("subclass with id '%v' not found - %v", c.Subclass.ID, err)
		}

		if subclass.ClassID != c.ClassID {
			return fmt.Errorf("can not assign subclass (id: %v) with class id %v to character (id: %v) with class id %v", subclass.ID, subclass.ClassID, c.ID, c.ClassID)
		}
	}

	var background Background
	err = db.Where("name = ?", c.BackgroundName).Find(&background).Error
	if err != nil {
		return fmt.Errorf("background with name '%v' not found - %v", c.BackgroundName, err)
	}

	if !utils.SliceContains(alignments, c.Alignment) {
		return fmt.Errorf("invalid alignment '%v'", c.Alignment)
	}
	if !utils.SliceContains(sizes, c.Size) {
		return fmt.Errorf("invalid size '%v'", c.Size)
	}

	err = validateStats(c)
	if err != nil {
		return errors.New(fmt.Sprintf("stats validation failed: %v", err))
	}

	return nil
}

func validateStats(c *Character) error {
	if !isValidStat(c.Strength) {
		return errors.New("invalid Strength")
	}
	if !isValidStat(c.Dexterity) {
		return errors.New("invalid Dexterity")
	}
	if !isValidStat(c.Constitution) {
		return errors.New("invalid Constitution")
	}
	if !isValidStat(c.Intelligence) {
		return errors.New("invalid Intelligence")
	}
	if !isValidStat(c.Wisdom) {
		return errors.New("invalid Wisdom")
	}
	if !isValidStat(c.Charisma) {
		return errors.New("invalid Charisma")
	}

	return nil
}

func isValidStat(stat int) bool {
	return stat >= 1 && stat <= 20
}
