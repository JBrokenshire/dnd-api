package models

type Race struct {
	ID               uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	CreatureType     string `json:"creature_type"`
	Size             string `json:"size"`
	BaseSpeed        int    `json:"base_speed"`

	Logo File `json:"logo"`
}
