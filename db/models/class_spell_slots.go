package models

type ClassSpellSlots struct {
	ID         uint `gorm:"primary_key;auto_increment" json:"id"`
	ClassId    uint `json:"class_id"`
	ClassLevel int  `json:"level"`
	SpellLevel int  `json:"spell_level"`
	SpellSlots int  `json:"spell_slots"`
}
