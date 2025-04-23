package models

type Class struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`

	Image File `json:"image"`
}
