package models

type Background struct {
	Name        string `gorm:"primary_key" json:"name"`
	Feature     string `json:"feature"`
	Description string `json:"description"`
}
