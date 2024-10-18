package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharactersLanguages() {
	charactersLanguages := []models.CharacterLanguage{
		{
			CharacterID: 1,
			Language:    "Gnomish",
		},
		{
			CharacterID: 1,
			Language:    "Halfling",
		},
		{
			CharacterID: 2,
			Language:    "Dwarvish",
		},
		{
			CharacterID: 3,
			Language:    "Giant",
		},
		{
			CharacterID: 3,
			Language:    "Orc",
		},
		{
			CharacterID: 4,
			Language:    "Draconic",
		},
		{
			CharacterID: 5,
			Language:    "Deep Speech",
		},
		{
			CharacterID: 5,
			Language:    "Halfling",
		},
		{
			CharacterID: 5,
			Language:    "Infernal",
		},
	}

	for _, language := range charactersLanguages {
		err := s.DB.Where("character_id = ? AND language = ?", language.CharacterID, language.Language).FirstOrCreate(&language).Error
		if err != nil {
			log.Printf("error creating character language for CharacterID: %q, Language: %s: %s", language.CharacterID, language.Language, err.Error())
		}
	}
}
