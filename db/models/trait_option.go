package models

type TraitOption struct {
	ID          uint `gorm:"primary_key;auto_increment"`
	TraitId     string
	Name        string
	Description string
}
