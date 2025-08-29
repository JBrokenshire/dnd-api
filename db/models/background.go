package models

type Background struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `json:"name"`
	Feature     string `json:"feature"`
	Description string `json:"description"`
}
