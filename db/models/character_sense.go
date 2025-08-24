package models

type CharacterSense struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	CharacterID uint   `json:"character_id"`
	Sense       string `json:"sense"`
}
