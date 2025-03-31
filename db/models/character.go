package models

type Character struct {
	ID      int    `json:"-" gorm:"primary_key"`
	Name    string `json:"name"`
	RaceId  int    `json:"race_id"`
	ClassId int    `json:"class_id"`

	Race  Race
	Class Class
}
