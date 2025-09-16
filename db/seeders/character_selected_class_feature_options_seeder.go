package seeders

import (
	m "dnd-api/db/models"
	"github.com/labstack/gommon/log"
)

func (s *Seeder) SetCharacterSelectedClassFeatureOptions() {
	characterSelectedClassFeatureOptions := []*m.CharacterSelectedClassFeatureOption{
		{
			ID:                   1,
			CharacterId:          1,
			ClassFeatureOptionId: 13,
		},
		{
			ID:                   2,
			CharacterId:          1,
			ClassFeatureOptionId: 25,
		},
		{
			ID:                   3,
			CharacterId:          1,
			ClassFeatureOptionId: 22,
		},
		{
			ID:                   4,
			CharacterId:          1,
			ClassFeatureOptionId: 26,
		},
		{
			ID:                   5,
			CharacterId:          1,
			ClassFeatureOptionId: 39,
		},
		{
			ID:                   6,
			CharacterId:          1,
			ClassFeatureOptionId: 30,
		},
		{
			ID:                   7,
			CharacterId:          1,
			ClassFeatureOptionId: 56,
		},
		{
			ID:                   8,
			CharacterId:          1,
			ClassFeatureOptionId: 57,
		},
		{
			ID:                   9,
			CharacterId:          1,
			ClassFeatureOptionId: 62,
		},
		{
			ID:                   10,
			CharacterId:          1,
			ClassFeatureOptionId: 63,
		},
	}

	for _, characterSelectedClassFeatureOption := range characterSelectedClassFeatureOptions {
		err := s.DB.Where("id = ?", characterSelectedClassFeatureOption.ID).FirstOrCreate(&characterSelectedClassFeatureOption).Error
		if err != nil {
			log.Printf("Error creating character selected class feature option with id %v in seeder: %v", characterSelectedClassFeatureOption.ID, err.Error())
		}
	}
}
