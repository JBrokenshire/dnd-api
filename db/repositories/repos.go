package repositories

import "github.com/jinzhu/gorm"

type Repos struct {
	DB        *gorm.DB
	User      *UserRepository
	Class     *ClassRepository
	Race      *RaceRepository
	Character *CharacterRepository
	File      *FileRepository
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		DB:        db,
		User:      NewUserRepository(db),
		Class:     NewClassRepository(db),
		Race:      NewRaceRepository(db),
		Character: NewCharacterRepository(db),
		File:      NewFileRepository(db),
	}
}
