package models

type Class struct {
	ID               uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	PrimaryAbility   string `json:"primary_ability"`
	HitPointDieValue int    `json:"hit_point_die_value"`
	Saves            string `json:"saves"`

	Logo File `json:"logo"`
}
