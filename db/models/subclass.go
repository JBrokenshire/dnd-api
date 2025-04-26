package models

type Subclass struct {
	ID               uint   `gorm:"primary_key" json:"id"`
	ClassId          uint   `json:"class_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`

	Logo File `json:"logo"`
}
