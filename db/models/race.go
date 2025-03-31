package models

type Race struct {
	ID   int    `json:"-" gorm:"primary_key"`
	Name string `json:"name"`
}
