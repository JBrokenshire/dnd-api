package models

type ClassSpellLevel struct {
	ID            uint `gorm:"primary_key;auto_increment"`
	ClassId       uint
	ClassLevel    uint
	SpellLevel    int
	NumberOfSlots int
}
