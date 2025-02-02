package repositories

import "github.com/jinzhu/gorm"

type Repos struct {
	DB    *gorm.DB
	Class *ClassRepository
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		DB:    db,
		Class: NewClassRepository(db),
	}
}
