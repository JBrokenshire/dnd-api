package models

type ClassFeature struct {
	ID          string `gorm:"primary_key"`
	ClassId     uint
	Level       int
	Name        string
	Description string
	Priority    int

	Options []*ClassFeatureOption
}
