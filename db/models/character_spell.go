package models

type CharacterSpell struct {
	ID          uint `gorm:"primary_key;auto_increment" json:"id"`
	CharacterID uint `json:"character_id"`
	SpellID     uint `json:"spell_id"`
}
