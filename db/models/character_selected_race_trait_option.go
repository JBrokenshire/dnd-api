package models

type CharacterSelectedRaceTraitOption struct {
	ID            uint `gorm:"primary_key;auto_increment"`
	CharacterId   uint
	TraitOptionId uint
}
