package models

type Class struct {
	ID               uint   `gorm:"primary_key" example:"1"`
	Name             string `example:"Barbarian"`
	ShortDescription string `example:"A short description"`
	LongDescription  string `example:"A longer description"`
}
