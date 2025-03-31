package responses

import m "dnd-api/db/models"

type CharacterResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Race  RaceResponse  `json:"race"`
	Class ClassResponse `json:"class"`
}

type CharacterPaginatedResponse struct {
	Data []CharacterResponse `json:"data"`
	Meta ResponseMeta        `json:"meta"`
}

func NewCharacterResponse(character *m.Character) *CharacterResponse {
	res := &CharacterResponse{
		Name: character.Name,
	}

	if character.Race.ID != 0 {
		res.Race = *NewRaceResponse(&character.Race)
	}
	if character.Class.ID != 0 {
		res.Class = *NewClassResponse(&character.Class)
	}

	return res
}

func NewCharacterResponses(characters []m.Character) []CharacterResponse {
	res := make([]CharacterResponse, 0)
	for _, character := range characters {
		res = append(res, *NewCharacterResponse(&character))
	}

	return res
}

func NewCharacterPaginatedResponse(characters []m.Character, count, page, pageSize int) *CharacterPaginatedResponse {
	return &CharacterPaginatedResponse{
		Data: NewCharacterResponses(characters),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
