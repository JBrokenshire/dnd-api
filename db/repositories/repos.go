package repositories

import (
	"github.com/jinzhu/gorm"
)

type Repos struct {
	DB                     *gorm.DB
	User                   *UserRepository
	Class                  *ClassRepository
	Subclass               *SubclassRepository
	Race                   *RaceRepository
	File                   *FileRepository
	Character              *CharacterRepository
	CharacterInventoryItem *CharacterInventoryItemRepository
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		DB:                     db,
		User:                   NewUserRepository(db),
		Class:                  NewClassRepository(db),
		Subclass:               NewSubclassRepository(db),
		Race:                   NewRaceRepository(db),
		File:                   NewFileRepository(db),
		Character:              NewCharacterRepository(db),
		CharacterInventoryItem: NewCharacterInventoryItemRepository(db),
	}
}
