package models

type Class struct {
	ID   int    `json:"-" gorm:"primary_key"`
	Name string `json:"name"`
}
