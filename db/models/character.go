package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Character struct {
	ID                int    `gorm:"autoIncrement;primary_key" json:"id"`
	Name              string `gorm:"not null" json:"name"`
	Level             int    `gorm:"not null" json:"level"`
	ProfilePictureURL string `json:"profile_picture_url"`
	ClassID           int    `json:"class_id"`
	RaceID            int    `json:"race_id"`

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

	AttacksPerAction int `gorm:"not null" json:"attacks_per_action"`

	Class Class `json:"class"`
	Race  Race  `json:"race"`
}

func (c *Character) BeforeCreate(_ *gorm.DB) error {
	err := validateStats(c)
	if err != nil {
		return err
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
