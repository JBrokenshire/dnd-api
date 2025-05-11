package models

type Character struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	ClassId  uint   `json:"class_id"`
	RaceId   uint   `json:"race_id"`
	Level    int    `json:"level"`
	Pronouns string `json:"pronouns"`

	// Ability Scores
	Strength     uint `json:"strength"`
	Dexterity    uint `json:"dexterity"`
	Constitution uint `json:"constitution"`
	Intelligence uint `json:"intelligence"`
	Wisdom       uint `json:"wisdom"`
	Charisma     uint `json:"charisma"`

	Class          Class `json:"class"`
	Race           Race  `json:"race"`
	ProfilePicture File  `json:"profile_picture"`
}
