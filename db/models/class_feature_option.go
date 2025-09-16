package models

type ClassFeatureOption struct {
	ID             uint `gorm:"primary_key;auto_increment"`
	ClassFeatureId string
	Name           string
	Description    string
}
