package models

import "time"

const (
	FileModelClassLogo               = "ClassLogo"
	FileModelClassBackgroundImage    = "ClassBackgroundImage"
	FileModelSubclassLogo            = "SubclassLogo"
	FileModelRaceLogo                = "RaceLogo"
	FileModelCharacterProfilePicture = "CharacterProfilePicture"
)

type File struct {
	ID           uint       `json:"-" gorm:"primary_key"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
	Model        string     `json:"model"`
	ModelId      uint       `json:"model_id"`
	Filename     string     `json:"filename"`
	FileLocation string     `json:"file_location"`
}
