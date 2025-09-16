package models

type CharacterSelectedClassFeatureOption struct {
	ID                   uint `gorm:"primary_key;auto_increment"`
	CharacterId          uint
	ClassFeatureOptionId uint
}
