package models

type RaceTrait struct {
	ID      uint `gorm:"primary_key;auto_increment"`
	RaceId  uint
	TraitId string
}
