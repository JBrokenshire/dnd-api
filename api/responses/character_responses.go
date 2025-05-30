package responses

import (
	m "dnd-api/db/models"
)

type CharacterResponse struct {
	ID       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	ClassId  uint   `json:"class_id"`
	RaceId   uint   `json:"race_id"`
	Pronouns string `json:"pronouns"`
	Level    int    `json:"level"`

	// Ability Scores
	Strength     uint `json:"strength"`
	Dexterity    uint `json:"dexterity"`
	Constitution uint `json:"constitution"`
	Intelligence uint `json:"intelligence"`
	Wisdom       uint `json:"wisdom"`
	Charisma     uint `json:"charisma"`

	Inspiration bool `json:"inspiration"`

	// Settings
	AdvancementType string `json:"advancement_type"`
	HitPointType    string `json:"hit_point_type"`

	Class          ClassResponse `json:"class"`
	Race           RaceResponse  `json:"race"`
	ProfilePicture FileResponse  `json:"profile_picture"`
}

type SlimCharacterResponse struct {
	ID       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	ClassId  uint   `json:"class_id"`
	RaceId   uint   `json:"race_id"`
	Pronouns string `json:"pronouns"`
	Level    int    `json:"level"`

	Class          ClassResponse `json:"class"`
	Race           RaceResponse  `json:"race"`
	ProfilePicture FileResponse  `json:"profile_picture"`
}

type CharacterPaginatedResponse struct {
	Data []SlimCharacterResponse `json:"data"`
	Meta ResponseMeta            `json:"meta"`
}

func NewCharacterResponse(character *m.Character) *CharacterResponse {
	res := &CharacterResponse{
		ID:       character.ID,
		UserId:   character.UserId,
		Name:     character.Name,
		ClassId:  character.ClassId,
		RaceId:   character.RaceId,
		Pronouns: character.Pronouns,
		Level:    character.Level,

		Strength:     character.Strength,
		Dexterity:    character.Dexterity,
		Constitution: character.Constitution,
		Intelligence: character.Intelligence,
		Wisdom:       character.Wisdom,
		Charisma:     character.Charisma,

		Inspiration: character.Inspiration,

		AdvancementType: character.AdvancementType,
		HitPointType:    character.HitPointType,
	}

	if character.Class.ID != 0 {
		res.Class = *NewClassResponse(&character.Class)
	}
	if character.Race.ID != 0 {
		res.Race = *NewRaceResponse(&character.Race)
	}
	if character.ProfilePicture.ID != 0 {
		res.ProfilePicture = *NewFileResponse(&character.ProfilePicture)
	}

	return res
}

func NewSlimCharacterResponse(character *m.Character) *SlimCharacterResponse {
	res := &SlimCharacterResponse{
		ID:       character.ID,
		UserId:   character.UserId,
		Name:     character.Name,
		ClassId:  character.ClassId,
		RaceId:   character.RaceId,
		Pronouns: character.Pronouns,
		Level:    character.Level,
	}

	if character.Class.ID != 0 {
		res.Class = *NewClassResponse(&character.Class)
	}
	if character.Race.ID != 0 {
		res.Race = *NewRaceResponse(&character.Race)
	}
	if character.ProfilePicture.ID != 0 {
		res.ProfilePicture = *NewFileResponse(&character.ProfilePicture)
	}

	return res
}

func NewCharacterResponses(characters []*m.Character) []SlimCharacterResponse {
	var res []SlimCharacterResponse
	for _, character := range characters {
		res = append(res, *NewSlimCharacterResponse(character))
	}
	return res
}

func NewCharacterPaginatedResponse(characters []*m.Character, count, page, pageSize int) *CharacterPaginatedResponse {
	return &CharacterPaginatedResponse{
		Data: NewCharacterResponses(characters),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
