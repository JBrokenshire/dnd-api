package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetClassFeatureOptions() {
	classFeatureOptions := []*m.ClassFeatureOption{
		{
			ID:             1,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Bagpipes",
		},
		{
			ID:             2,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Birdpipes",
		},
		{
			ID:             3,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Clarinet",
		},
		{
			ID:             4,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Drum",
		},
		{
			ID:             5,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Dulcimer",
		},
		{
			ID:             6,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Fiddle",
		},
		{
			ID:             7,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Flute",
		},
		{
			ID:             8,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Glaur",
		},
		{
			ID:             9,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Hand Drum",
		},
		{
			ID:             10,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Harp",
		},
		{
			ID:             11,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Horn",
		},
		{
			ID:             12,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Longhorn",
		},
		{
			ID:             13,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Lute",
		},
		{
			ID:             14,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Lyre",
		},
		{
			ID:             15,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Pan Flute",
		},
		{
			ID:             16,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Shawm",
		},
		{
			ID:             17,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Songhorn",
		},
		{
			ID:             18,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Tantan",
		},
		{
			ID:             19,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Thelarr",
		},
		{
			ID:             20,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Tocken",
		},
		{
			ID:             21,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Trumpet",
		},
		{
			ID:             22,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Viol",
		},
		{
			ID:             23,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Wargong",
		},
		{
			ID:             24,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Yarting",
		},
		{
			ID:             25,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Zulkoon",
		},
		{
			ID:             26,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Acrobatics",
		},
		{
			ID:             27,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Animal Handling",
		},
		{
			ID:             28,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Arcana",
		},
		{
			ID:             29,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Athletics",
		},
		{
			ID:             30,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Deception",
		},
		{
			ID:             31,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "History",
		},
		{
			ID:             32,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Insight",
		},
		{
			ID:             33,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Intimidation",
		},
		{
			ID:             34,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Investigation",
		},
		{
			ID:             35,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Medicine",
		},
		{
			ID:             36,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Nature",
		},
		{
			ID:             37,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Perception",
		},
		{
			ID:             38,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Performance",
		},
		{
			ID:             39,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Persuasion",
		},
		{
			ID:             40,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Religion",
		},
		{
			ID:             41,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Sleight of Hand",
		},
		{
			ID:             42,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Stealth",
		},
		{
			ID:             43,
			ClassFeatureId: "bard-core-bard-traits",
			Name:           "Survival",
		},
		{
			ID:             44,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Acrobatics",
		},
		{
			ID:             45,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Animal Handling",
		},
		{
			ID:             46,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Arcana",
		},
		{
			ID:             47,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Athletics",
		},
		{
			ID:             48,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Deception",
		},
		{
			ID:             49,
			ClassFeatureId: "bard-expertise-2",
			Name:           "History",
		},
		{
			ID:             50,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Insight",
		},
		{
			ID:             51,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Intimidation",
		},
		{
			ID:             52,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Investigation",
		},
		{
			ID:             53,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Medicine",
		},
		{
			ID:             54,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Nature",
		},
		{
			ID:             55,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Perception",
		},
		{
			ID:             56,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Performance",
		},
		{
			ID:             57,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Persuasion",
		},
		{
			ID:             58,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Religion",
		},
		{
			ID:             59,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Sleight of Hand",
		},
		{
			ID:             60,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Stealth",
		},
		{
			ID:             61,
			ClassFeatureId: "bard-expertise-2",
			Name:           "Survival",
		},
		{
			ID:             62,
			ClassFeatureId: "bard-subclass",
			Name:           "College of Glamour",
		},
		{
			ID:             63,
			ClassFeatureId: "bard-asi-4",
			Name:           "Feat",
			Description:    "Fey Touched",
		},
		{
			ID:             63,
			ClassFeatureId: "bard-asi-4",
			Name:           "Ability Score Improvement",
		},
	}

	for _, classFeatureOption := range classFeatureOptions {
		err := s.DB.Where("id = ?", classFeatureOption.ID).FirstOrCreate(&classFeatureOption).Error
		if err != nil {
			log.Printf("Error creating class feature option with id %v in seeder: %v", classFeatureOption.ID, err.Error())
		}
	}
}
