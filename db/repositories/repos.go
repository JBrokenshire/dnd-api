package repositories

import "github.com/jinzhu/gorm"

type Repos struct {
	DB    *gorm.DB
	User  *UserRepository
	Class *ClassRepository
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		DB:    db,
		User:  NewUserRepository(db),
		Class: NewClassRepository(db),
	}
}
