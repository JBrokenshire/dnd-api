package models

type CharacterSpell struct {
	ID          int `gorm:"primary_key" json:"id"`
	CharacterID int `json:"character_id"`
	SpellID     int `json:"spell_id"`

	Spell Spell `json:"spell"`
}
